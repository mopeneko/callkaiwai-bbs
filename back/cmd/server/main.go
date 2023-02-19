package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/mopeneko/callkaiwai-bbs/back/ent"
	"github.com/mopeneko/callkaiwai-bbs/back/gen/callkaiwaibbs/v1/callkaiwaibbsv1connect"
	"github.com/mopeneko/callkaiwai-bbs/back/handler/callkaiwaiservice"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/xerrors"
)

func main() {
	db, err := initDB(context.Background())
	if err != nil {
		log.Fatalf("failed to initialize database: %+v", err)
	}

	defer db.Close()

	service := callkaiwaiservice.NewCallKaiwaiServiceHandler(db)
	mux := http.NewServeMux()
	path, handler := callkaiwaibbsv1connect.NewCallKaiwaiBBSServiceHandler(service)
	mux.Handle(path, handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	log.Println("Listening on", port)

	http.ListenAndServe(port, h2c.NewHandler(mux, &http2.Server{}))
}

func initDB(ctx context.Context) (*ent.Client, error) {
	cfg := mysql.Config{
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST"),
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := ent.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, xerrors.Errorf("failed to open database: %w", err)
	}

	if err := db.Schema.Create(ctx); err != nil {
		db.Close()
		return nil, xerrors.Errorf("failed to create schema: %w", err)
	}

	return db, nil
}
