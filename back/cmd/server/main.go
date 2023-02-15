package main

import (
	"context"
	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	callkaiwaibbsv1 "github.com/mopeneko/callkaiwai-bbs/back/gen/callkaiwaibbs/v1"
	"github.com/mopeneko/callkaiwai-bbs/back/gen/callkaiwaibbs/v1/callkaiwaibbsv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type CallKaiwaiService struct{}

func (s *CallKaiwaiService) Ping(ctx context.Context, _ *connect.Request[callkaiwaibbsv1.PingRequest]) (*connect.Response[callkaiwaibbsv1.PingResponse], error) {
	return connect.NewResponse(&callkaiwaibbsv1.PingResponse{
		Text: "Pong",
	}), nil
}

func main() {
	service := &CallKaiwaiService{}
	mux := http.NewServeMux()
	path, handler := callkaiwaibbsv1connect.NewCallKaiwaiBBSServiceHandler(service)
	mux.Handle(path, handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	http.ListenAndServe(port, h2c.NewHandler(mux, &http2.Server{}))
}
