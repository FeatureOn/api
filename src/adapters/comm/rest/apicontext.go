package rest

import (
	"fmt"
	"net/http"
	"time"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/application"

	middleware "dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest/middleware"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	openapimw "github.com/go-openapi/runtime/middleware"

	"github.com/uber/jaeger-client-go"
	jprom "github.com/uber/jaeger-lib/metrics/prometheus"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

// DBContext is the interface that APIContext expects to be fulfilled by Database code
// type DBContext interface {
// 	CheckConnection() bool
// }

// APIContext handler for getting and updating Ratings
type APIContext struct {
	validation *middleware.Validation
	//dbContext  DBContext
	healthRepo application.HealthRepository
	userRepo   application.UserRepository
}

// NewAPIContext returns a new APIContext handler with the given logger
//func NewAPIContext(dc DBContext, bindAddress *string, ur application.UserRepository) *http.Server {
func NewAPIContext(bindAddress *string, hr application.HealthRepository, ur application.UserRepository) *http.Server {
	apiContext := &APIContext{
		healthRepo: hr,
		userRepo:   ur,
	}
	s := apiContext.prepareContext(bindAddress)
	return s
}

func (apiContext *APIContext) prepareContext(bindAddress *string) *http.Server {
	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := jprom.New()

	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg, err := jaegercfg.FromEnv()
	if err != nil || cfg.ServiceName == "" {
		cfg = &jaegercfg.Configuration{
			ServiceName: "GoBoiler.WebApi",
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans: true,
			},
		}
	}

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, _ := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	env.Parse()

	//v := middleware.NewValidation()

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()
	sm.Use(middleware.MetricsMiddleware)

	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/", apiContext.Index)
	getR.HandleFunc("/version", apiContext.Version)
	getR.HandleFunc("/health/live", apiContext.Live)
	getR.HandleFunc("/health/ready", apiContext.Ready)
	getR.HandleFunc("/user/{id}", apiContext.GetUser)

	// handler for documentation
	opts := openapimw.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := openapimw.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// create a new server
	s := &http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      sm,                // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	sm.PathPrefix("/metrics").Handler(promhttp.Handler())
	prometheus.MustRegister(middleware.RequestCounterVec)
	prometheus.MustRegister(middleware.RequestDurationGauge)

	return s
}

// createSpan creates a new openTracing.Span with the given name and returns it
func createSpan(spanName string, r *http.Request) (span opentracing.Span) {
	tracer := opentracing.GlobalTracer()

	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		// The method is called without a span context in the http header.
		//
		span = tracer.StartSpan(spanName)
	} else {
		// Create the span referring to the RPC client if available.
		// If wireContext == nil, a root span will be created.
		span = opentracing.StartSpan(spanName, ext.RPCServerOption(wireContext))
	}
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, r.URL.RequestURI())
	ext.HTTPMethod.Set(span, r.Method)
	return span
}

// ErrInvalidRatingPath is an error message when the Rating path is not valid
var ErrInvalidRatingPath = fmt.Errorf("Invalid Path, path should be /Details/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}
