package http

import (
	"net/http"
)

func newFileServer() *http.ServeMux {
	fileshttp := http.NewServeMux()
	fileshttp.Handle("/", http.FileServer(http.Dir("./web/public/")))
	return fileshttp
}
