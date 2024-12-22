package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"snippetbox/internal/models"

	_ "github.com/lib/pq"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	dsn := flag.String("dsn", `postgresql://<username>:<password>@<host>:<port>/<database>?<parameters>`, "PostgreSQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
	}
	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	mux := app.routes()

	logger.Info("Starting server", "addr", *addr)
	err = http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
