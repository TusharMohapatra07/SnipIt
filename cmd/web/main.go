package main

import (
	"database/sql"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"snippetbox/internal/models"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	logger         *slog.Logger
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	sessionManager *scs.SessionManager
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error(err.Error())
	}
	connStr := os.Getenv("CONNSTR")
	addr := os.Getenv("ADDRESS")

	db, err := openDB(connStr)
	if err != nil {
		logger.Error(err.Error())
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(db)
	sessionManager.Lifetime = 9 * time.Hour

	app := &application{
		logger:         logger,
		snippets:       &models.SnippetModel{DB: db},
		templateCache:  templateCache,
		sessionManager: sessionManager,
	}

	mux := app.routes()

	logger.Info("Starting server", "addr", addr)
	err = http.ListenAndServe(addr, mux)

	logger.Error(err.Error())
	os.Exit(1)
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
