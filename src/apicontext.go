package net

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// APIContext handler for getting and updating Ratings
type APIContext struct {
	v *Validation
}

// DBContext is the struct that has a MongoDB connection together with standard APIContext. It's used for handler functions which will use database
type DBContext struct {
	MongoClient  mongo.Client
	DatabaseName string
	APIContext
}

// NewAPIContext returns a new APIContext handler with the given logger
func NewAPIContext(v *Validation) *APIContext {
	return &APIContext{v}
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
