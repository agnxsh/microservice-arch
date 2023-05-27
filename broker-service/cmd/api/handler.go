package main

import (
	"encoding/json"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:	false,
		Message:	"Ping the broker",
	}

	out,_ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)

}