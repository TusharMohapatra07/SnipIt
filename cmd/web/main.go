package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	app := &application{
		logger: logger,
	}

	mux := app.routes()

	logger.Info("Starting server", "addr", *addr)
	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())
}
