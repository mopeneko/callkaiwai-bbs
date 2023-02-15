// Code generated by protoc-gen-connect-go.exe. DO NOT EDIT.
//
// Source: callkaiwaibbs/v1/callkaiwai.proto

package callkaiwaibbsv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/mopeneko/callkaiwai-bbs/back/gen/callkaiwaibbs/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// CallKaiwaiBBSServiceName is the fully-qualified name of the CallKaiwaiBBSService service.
	CallKaiwaiBBSServiceName = "callkaiwaibbs.v1.CallKaiwaiBBSService"
)

// CallKaiwaiBBSServiceClient is a client for the callkaiwaibbs.v1.CallKaiwaiBBSService service.
type CallKaiwaiBBSServiceClient interface {
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
}

// NewCallKaiwaiBBSServiceClient constructs a client for the callkaiwaibbs.v1.CallKaiwaiBBSService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCallKaiwaiBBSServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CallKaiwaiBBSServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &callKaiwaiBBSServiceClient{
		ping: connect_go.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+"/callkaiwaibbs.v1.CallKaiwaiBBSService/Ping",
			opts...,
		),
	}
}

// callKaiwaiBBSServiceClient implements CallKaiwaiBBSServiceClient.
type callKaiwaiBBSServiceClient struct {
	ping *connect_go.Client[v1.PingRequest, v1.PingResponse]
}

// Ping calls callkaiwaibbs.v1.CallKaiwaiBBSService.Ping.
func (c *callKaiwaiBBSServiceClient) Ping(ctx context.Context, req *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// CallKaiwaiBBSServiceHandler is an implementation of the callkaiwaibbs.v1.CallKaiwaiBBSService
// service.
type CallKaiwaiBBSServiceHandler interface {
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
}

// NewCallKaiwaiBBSServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCallKaiwaiBBSServiceHandler(svc CallKaiwaiBBSServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/callkaiwaibbs.v1.CallKaiwaiBBSService/Ping", connect_go.NewUnaryHandler(
		"/callkaiwaibbs.v1.CallKaiwaiBBSService/Ping",
		svc.Ping,
		opts...,
	))
	return "/callkaiwaibbs.v1.CallKaiwaiBBSService/", mux
}

// UnimplementedCallKaiwaiBBSServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCallKaiwaiBBSServiceHandler struct{}

func (UnimplementedCallKaiwaiBBSServiceHandler) Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("callkaiwaibbs.v1.CallKaiwaiBBSService.Ping is not implemented"))
}
