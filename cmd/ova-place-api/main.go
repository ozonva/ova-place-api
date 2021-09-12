package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/ozonva/ova-place-api/internal/api"
	flusher "github.com/ozonva/ova-place-api/internal/flusher"
	"github.com/ozonva/ova-place-api/internal/health"
	"github.com/ozonva/ova-place-api/internal/metrics"
	"github.com/ozonva/ova-place-api/internal/producer"
	"github.com/ozonva/ova-place-api/internal/repo"
	"github.com/ozonva/ova-place-api/internal/saver"
	"github.com/ozonva/ova-place-api/internal/tracing"
	desc "github.com/ozonva/ova-place-api/pkg/ova-place-api"
)

const (
	logFilePath   = "../logs/system.log"
	maxLogSize    = 10
	maxLogBackups = 5
	maxLogAge     = 28
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln(err)
	}
}

func runGrpc(producer *producer.Producer, listener net.Listener, loggerInstance zerolog.Logger, saverInstance saver.Saver, repoInstance repo.Repo) *grpc.Server {
	grpcServer := grpc.NewServer()

	cudCounterInstance := metrics.NewCudCounter(
		promauto.NewCounter(prometheus.CounterOpts{
			Name: "successful_creates",
			Help: "The total number of successful_creates",
		}),
		promauto.NewCounter(prometheus.CounterOpts{
			Name: "successful_updates",
			Help: "The total number of successful_updates",
		}),
		promauto.NewCounter(prometheus.CounterOpts{
			Name: "successful_deletes",
			Help: "The total number of successful_deletes",
		}),
	)

	desc.RegisterOvaPlaceApiV1Server(grpcServer, api.NewOvaPlaceAPI(repoInstance, saverInstance, *producer, cudCounterInstance, loggerInstance))

	go func(server *grpc.Server) {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}(grpcServer)

	return grpcServer
}

func runHTTPServers(db *sqlx.DB, producer *producer.Producer) []*http.Server {
	metricsMux := http.NewServeMux()
	metricsMux.Handle(os.Getenv("PROMETHEUS_PATH"), promhttp.Handler())
	metricsServer := http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv("PROMETHEUS_PORT")),
		Handler: metricsMux,
	}

	healthService := health.NewHealth(health.NewDatabaseCheck(db), health.NewKafkaCheck(producer))
	healthMux := http.NewServeMux()
	healthMux.HandleFunc(os.Getenv("HEALTH_PATH"), healthCheckHandler(healthService))
	healthServer := http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv("HEALTH_PORT")),
		Handler: healthMux,
	}

	go func(metricsServer *http.Server) {
		err := metricsServer.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to init metrics: %v", err)
		}
	}(&metricsServer)

	go func(healthServer *http.Server) {

		err := healthServer.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to init healthcheck: %v", err)
		}
	}(&healthServer)

	return []*http.Server{&metricsServer, &healthServer}
}

func healthCheckHandler(healthService health.Health) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.MarshalIndent(healthService.Check(), "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write(bytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	lumberjackInstance := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    maxLogSize,
		MaxBackups: maxLogBackups,
		MaxAge:     maxLogAge,
		Compress:   true,
	}

	z := zerolog.New(lumberjackInstance)
	loggerInstance := z.With().Caller().Timestamp().Logger()

	db, err := sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_STRING"))
	if err != nil {
		log.Fatalln(err)
	}

	producerInstance, err := producer.NewProducer([]string{os.Getenv("KAFKA_BROKER_URL")})
	if err != nil {
		log.Fatalf("failed to init the producer: %v", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repoInstance := repo.NewRepo(db)
	flusherInstance := flusher.NewFlusher(getIntEnv("BATCH_SIZE"), repoInstance)
	saverInstance := saver.NewSaver(context.Background(), uint(getIntEnv("SAVER_CAPACITY")), time.Second*time.Duration(getIntEnv("SAVER_TIMEOUT")), flusherInstance)

	tracer, closer := tracing.Init("ova-place-api")
	opentracing.SetGlobalTracer(tracer)

	httpServers := runHTTPServers(db, &producerInstance)
	grpcServer := runGrpc(&producerInstance, listener, loggerInstance, saverInstance, repoInstance)

	defer freeUpResources(httpServers, grpcServer, db, producerInstance, listener, closer, lumberjackInstance, saverInstance)

	if <-sigCh; true {
		log.Println("Gracefully stopping")
		freeUpResources(httpServers, grpcServer, db, producerInstance, listener, closer, lumberjackInstance, saverInstance)
		os.Exit(0)
	}
}

func freeUpResources(httpServers []*http.Server,
	grpcServer *grpc.Server,
	db *sqlx.DB,
	producerInstance producer.Producer,
	listener net.Listener,
	closer io.Closer,
	lumberjackInstance *lumberjack.Logger,
	saverInstance saver.Saver) {

	for index := range httpServers {
		err := httpServers[index].Shutdown(context.Background())
		if err != nil {
			log.Fatalf("failed to close the httpServer: %v", err)
		}
	}

	grpcServer.GracefulStop()

	err := listener.Close()
	if err != nil {
		log.Fatalf("failed to close the listener: %v", err)
	}

	err = closer.Close()
	if err != nil {
		log.Fatalf("failed to close tracing: %v", err)
	}

	err = db.Close()
	if err != nil {
		log.Fatalf("failed to close the db: %v", err)
	}

	err = producerInstance.Close()
	if err != nil {
		log.Fatalf("failed to close the producer: %v", err)
	}

	err = lumberjackInstance.Close()
	if err != nil {
		log.Fatalf("failed to close the lumberjack: %v", err)
	}

	err = saverInstance.Close()
	if err != nil {
		log.Fatalf("failed to close the saver: %v", err)
	}
}

func getIntEnv(key string) int {
	val := os.Getenv(key)
	ret, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("cannot parse int value: %v", err)
	}
	return ret
}
