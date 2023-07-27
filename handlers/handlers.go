package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var lock = make(chan struct{}, 1)

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CommandHandler(w http.ResponseWriter, req *http.Request) {
	select {
	case lock <- struct{}{}:
		defer func() { <-lock }() // Unlock deferred
	default:
		http.Error(w, "command already processing", 503)
		return
	}

	var cmd Command
	err := json.NewDecoder(req.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, "could not read json body", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	cmdOut, err := runCmd(cmd)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed running cmd: %+v", err), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("stdout:\n" + cmdOut.Stdout + "stderr:\n" + cmdOut.Stderr))
	if err != nil {
		http.Error(w, "could not write response", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}
