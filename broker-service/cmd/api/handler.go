package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action	string			`json:"action"`
	Auth 	AuthPayload		`json:"auth,omitpay"`
}

type AuthPayload struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:	false,
		Message:	"Ping the broker",
	}

	//calling a helper function instead of writing excessive lines of code :)
	_ = app.writeJSON(w, http.StatusOK, payload)

	// out,_ := json.MarshalIndent(payload, "", "\t")
	// w.Header().Set("Content-Type","application/json")
	// w.WriteHeader(http.StatusAccepted)
	// w.Write(out)

}

func (app *Config) HandleSubmission (w http.ResponseWriter, r *http.Request){
	 var requestPayload RequestPayload

	 err := app.readJSON(w, r, &requestPayload)
	 if err != nil {
		app.errorJSON(w, err)
		return
	 }

	 switch requestPayload.Action {
	
	 case "auth":
	
	 default:
		app.errorJSON(w, errors.New("unknown action"))
	 }
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload){
	//creating some json to the send to the auth microservice
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	//call the microservice
	request, err := http.NewRequest("POST","http://authentication-service/authenticate",bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	//make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("Error calling auth service"))
		return
	}

}