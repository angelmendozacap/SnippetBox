package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(contextKeyIsAuthenticated).(bool)
	if !ok {
		return false
	}
	return isAuthenticated
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) displayMessage(w http.ResponseWriter, code int, m interface{}) {
	j, err := json.Marshal(m)
	if err != nil {
		app.errorLog.Output(2, fmt.Sprintf("Error al convertir el mensaje: %s", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.WriteHeader(code)
	w.Write(j)
}
