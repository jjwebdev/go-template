package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/blockninja/ninjarouter"
)

type funcDetails struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type errorResponse struct {
	Code    int         `json:"code"`
	Error   string      `json:"error"`
	Message funcDetails `json:"message"`
}

func handleErrorAndRespond(details *funcDetails, w http.ResponseWriter, errString string, code int) {
	log.Println(details.Name, "handler error:", errString)
	w.WriteHeader(code)
	response := &errorResponse{
		Code:    code,
		Error:   errString,
		Message: *details,
	}
	json.NewEncoder(w).Encode(response)
}

func marshalAndRespond(details *funcDetails, w http.ResponseWriter, result interface{}) {
	response, err := json.Marshal(result)
	if err != nil {
		handleErrorAndRespond(details, w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(response)
	}
}

func mustGetID(r *http.Request, name string) int {
	nameStr := ninjarouter.Var(r, name)
	nameInt, err := strconv.Atoi(nameStr)
	if err != nil {
		panic(err)
	}
	return nameInt
}

func mustDecodeJSON(r *http.Request, target interface{}) {
	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		panic(err)
	}
}