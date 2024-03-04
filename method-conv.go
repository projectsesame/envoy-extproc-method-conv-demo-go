package main

import (
	"net/http"

	ep "github.com/wrossmorrow/envoy-extproc-sdk-go"
)

type methodConvRequestProcessor struct {
	opts *ep.ProcessingOptions
}

func (s *methodConvRequestProcessor) GetName() string {
	return "method-conv"
}

func (s *methodConvRequestProcessor) GetOptions() *ep.ProcessingOptions {
	return s.opts
}

func (s *methodConvRequestProcessor) ProcessRequestHeaders(ctx *ep.RequestContext, headers ep.AllHeaders) error {
	ok := false

	method := ""
	if ctx.Method == http.MethodGet {
		ok = true
		method = http.MethodPost
	}
	if ctx.Method == http.MethodPost {
		ok = true
		method = http.MethodGet
	}

	if ok {
		ctx.OverwriteHeader(":method", ep.HeaderValue{RawValue: []byte(method)})
	}
	return ctx.ContinueRequest()
}

func (s *methodConvRequestProcessor) ProcessRequestBody(ctx *ep.RequestContext, body []byte) error {
	return ctx.ContinueRequest()
}

func (s *methodConvRequestProcessor) ProcessRequestTrailers(ctx *ep.RequestContext, trailers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *methodConvRequestProcessor) ProcessResponseHeaders(ctx *ep.RequestContext, headers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *methodConvRequestProcessor) ProcessResponseBody(ctx *ep.RequestContext, body []byte) error {
	return ctx.ContinueRequest()
}

func (s *methodConvRequestProcessor) ProcessResponseTrailers(ctx *ep.RequestContext, trailers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *methodConvRequestProcessor) Init(opts *ep.ProcessingOptions, nonFlagArgs []string) error {
	s.opts = opts
	return nil
}

func (s *methodConvRequestProcessor) Finish() {}
