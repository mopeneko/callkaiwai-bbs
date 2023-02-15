package main

import (
	"net/http"
	"os"

	"github.com/mopeneko/callkaiwai-bbs/back/gen/callkaiwaibbs/v1/callkaiwaibbsv1connect"
	"github.com/mopeneko/callkaiwai-bbs/back/handler/callkaiwaiservice"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	service := callkaiwaiservice.NewCallKaiwaiServiceHandler()
	mux := http.NewServeMux()
	path, handler := callkaiwaibbsv1connect.NewCallKaiwaiBBSServiceHandler(service)
	mux.Handle(path, handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	http.ListenAndServe(port, h2c.NewHandler(mux, &http2.Server{}))
}
