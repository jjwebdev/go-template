package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"log"
)

type errorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type funcDetails struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func notFound(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "Not Found",
		Description: "This path does not exist",
	}
	log.Println(r.URL.EscapedPath())
	handleErrorAndRespond(details, w, "Path does not exist", http.StatusNotFound)
}

func marshalAndRespond(details *funcDetails, w http.ResponseWriter, result interface{}) {
	response, err := json.Marshal(result)
	if err != nil {
		handleErrorAndRespond(details, w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(response)

	}
}

func mustParseInt(r *http.Request, name string) int {
	result, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}
	return result

}

func mustDecodeJSON(r *http.Request, target interface{}) {
	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		panic(err)
	}
}
func handleErrorAndRespond(details *funcDetails, w http.ResponseWriter, errString string, code int) {
	log.Println("[HANDLER]", details.Name+":", errString)
	w.WriteHeader(code)
	response := &errorResponse{
		Code:  code,
		Error: errString,
	}
	json.NewEncoder(w).Encode(response)
}

func NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, urlStr, body)
}
