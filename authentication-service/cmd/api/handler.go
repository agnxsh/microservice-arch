package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate ( w http.ResponseWriter, r *http.Request){
	var requestPayload struct {

		Email		string		`json:"email"`
		Password	string		`json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err!=nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	//validating the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	//giving the same error message for security reasons
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest);
		return
	}

	//when the user logs in correctly
	payload := jsonResponse {
		Error: 	false,
		Message: fmt.Sprintf("Logged in user as %s", user.Email),
		Data: user,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
