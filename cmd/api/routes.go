package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

const baseURL = "/api"

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, app.authenticate)

	mux := pat.New()
	mux.Get(baseURL+"/snippets", dynamicMiddleware.ThenFunc(app.getSnippets))
	mux.Post(baseURL+"/snippets", dynamicMiddleware.ThenFunc(app.createSnippet))
	mux.Get(baseURL+"/snippets/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	return standardMiddleware.Then(mux)
}
