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
	dynamicMux.HandleFunc("GET /user/signup", app.userSignup)
	dynamicMux.HandleFunc("POST /user/signup", app.userSignupPost)
	dynamicMux.HandleFunc("GET /user/login", app.userLogin)
	dynamicMux.HandleFunc("POST /user/login", app.userLoginPost)
	dynamicMux.HandleFunc("GET /user/logout", app.userLogoutPost)

	mux.Handle("/", app.sessionManager.LoadAndSave(dynamicMux))

	return app.panicRecovery(app.logRequest(commonHeaders(mux)))
}
