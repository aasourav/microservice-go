package main

import (
	"net/http"

	"logger.svc/data"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestedPayload JSONPayload
	_ = app.readJSON(w, r, &requestedPayload)

	//insert dta
	event := data.LogEntry{
		Name: requestedPayload.Name,
		Data: requestedPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
