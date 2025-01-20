package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamicMux := http.NewServeMux()
	dynamicMux.HandleFunc("GET /{$}", app.home)
	dynamicMux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	dynamicMux.HandleFunc("GET /snippet/create", app.snippetCreate)
	dynamicMux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	mux.Handle("/", app.sessionManager.LoadAndSave(dynamicMux))

	return app.panicRecovery(app.logRequest(commonHeaders(mux)))
}
