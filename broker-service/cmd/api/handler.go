package main

import (
	"net/http"
)

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