package http

import (
	"net/http"
)

func NewFileServer() *http.ServeMux {
	fileshttp := http.NewServeMux()
	fileshttp.Handle("/", http.FileServer(http.Dir("./web/public/")))
	return fileshttp
}
