package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ozonva/ova-place-api/internal/metrics"
	"github.com/ozonva/ova-place-api/internal/producer"

	"github.com/opentracing/opentracing-go"

	flusher "github.com/ozonva/ova-place-api/internal/flusher"
	"github.com/ozonva/ova-place-api/internal/tracing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/ozonva/ova-place-api/internal/api"
	"github.com/ozonva/ova-place-api/internal/repo"
	desc "github.com/ozonva/ova-place-api/pkg/ova-place-api"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln(err)
	}
}

func runGrpc() error {
	grpcEndpoint := flag.String("grpc-server-endpoint", os.Getenv("GRPC_HOST")+os.Getenv("GRPC_PORT"), "gRPC server endpoint")

	db, err := sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_STRING"))
	if err != nil {
		log.Fatalln(err)
	}

	listen, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tracer, closer := tracing.Init("ova-place-api")
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			log.Fatalf("failed to init tracing: %v", err)
		}
	}(closer)
	opentracing.SetGlobalTracer(tracer)

	s := grpc.NewServer()
	repoInstance := repo.NewRepo(db)
	flusherInstance := flusher.NewFlusher(2, repoInstance)
	producerInstance, err := producer.NewProducer([]string{os.Getenv("KAFKA_BROKER_URL")})
	if err != nil {
		log.Fatalf("failed to init the producer: %v", err)
	}

	defer func(producerInstance producer.Producer) {
		err := producerInstance.Close()
		if err != nil {
			log.Fatalf("failed to close the producer: %v", err)
		}
	}(producerInstance)

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

	desc.RegisterOvaPlaceApiV1Server(s, api.NewOvaPlaceApi(repoInstance, flusherInstance, producerInstance, cudCounterInstance))

	log.Printf("Server listening on %s\n", *grpcEndpoint)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func runMetrics() {
	http.Handle(os.Getenv("PROMETHEUS_PATH"), promhttp.Handler())
	err := http.ListenAndServe(os.Getenv("PROMETHEUS_PORT"), nil)
	if err != nil {
		log.Fatalf("failed to init monitoring: %v", err)
	}
}

func main() {
	flag.Parse()

	go runMetrics()

	err := runGrpc()
	if err != nil {
		log.Fatalf("grpc not started")
	}
}
