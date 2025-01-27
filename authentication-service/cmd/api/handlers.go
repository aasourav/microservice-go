package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestedPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestedPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestedPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestedPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"))
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Loged in user %s", user.Email),
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
