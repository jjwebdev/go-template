package http

import (
	"log"
	"net/http"
	"time"
)

func publicChain(next http.HandlerFunc) http.HandlerFunc {
	return withLogging(http.HandlerFunc(next)).ServeHTTP
}

func userChain(next http.HandlerFunc) http.HandlerFunc {
	return withLogging(http.HandlerFunc(next)).ServeHTTP
}

func adminChain(next http.HandlerFunc) http.HandlerFunc {
	return withLogging(http.HandlerFunc(next)).ServeHTTP
}

func withLogging(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("-> [%s] %q\n", r.Method, r.URL.String())
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("<- [%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}