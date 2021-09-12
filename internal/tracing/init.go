package tracing

import (
	"io"
	"log"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	cfg.Reporter.LocalAgentHostPort = os.Getenv("JAEGER_AGENT_URL")

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatalf("failed to init Jaeger: %v", err)
	}
	return tracer, closer
}
