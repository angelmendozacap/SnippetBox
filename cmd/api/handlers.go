package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) getSnippets(w http.ResponseWriter, r *http.Request) {
	mr := ResponseMessage{}

	s, err := app.snippets.Latest()
	if err != nil {
		mr.AddError(
			http.StatusInternalServerError,
			"Cannot retrieve the information",
			"Check the service logs for details",
		)

		app.displayMessage(w, http.StatusInternalServerError, mr)
		return
	}

	mr.AddMessage(http.StatusOK, "Snippets retrieved successfully.", "")
	mr.Data = s

	app.displayMessage(w, http.StatusOK, mr)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	mr := ResponseMessage{}

	s, err := app.snippets.Get(id)
	if err != nil {
		mr.AddError(
			http.StatusInternalServerError,
			"Cannot retrieve the information",
			"Check the service logs for details",
		)

		app.displayMessage(w, http.StatusInternalServerError, mr)
		return
	}

	mr.AddMessage(http.StatusOK, "Snippet retrieved successfully.", "")
	mr.Data = s

	app.displayMessage(w, http.StatusOK, mr)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	mr := ResponseMessage{}

	body := struct {
		Title     string `json:"title"`
		Content   string `json:"content"`
		ExpiresAt string `json:"expires_at"`
		UserID    int    `json:"user_id"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		mr.AddError(
			http.StatusUnprocessableEntity,
			"Cannot read the sended data",
			"Check the service logs for details",
		)

		app.displayMessage(w, http.StatusUnprocessableEntity, mr)
		return
	}

	app.infoLog.Println("body", body)

	lastID, err := app.snippets.Insert(body.Title, body.Content, body.ExpiresAt, body.UserID)
	if err != nil {
		mr.AddError(
			http.StatusBadRequest,
			fmt.Sprintf("Cannot insert snippet: %s", err),
			"Check the service logs for details",
		)

		app.displayMessage(w, http.StatusBadRequest, mr)
		return
	}

	mr.AddMessage(http.StatusCreated, "Snippet created successfully.", fmt.Sprintf("/snippets/%v", lastID))

	app.displayMessage(w, http.StatusCreated, mr)
}
