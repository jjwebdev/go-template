package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func insecureChain(h http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return withLogging(
		withRecover(
			http.HandlerFunc(h),
		),
	).ServeHTTP
}

func withRecover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			rec := recover()
			if rec != nil {
				switch kind := rec.(type) {
				case string:
					err = errors.New(kind)
				case error:
					err = kind
				default:
					err = errors.New("Unknown error")
				}
				fmt.Println("panic: ", err.Error())
				errResp := errorResponse{
					Code:  http.StatusInternalServerError,
					Error: err.Error(),
				}
				result, err := json.Marshal(errResp)
				if err != nil {
					return
				}
				http.Error(w, string(result), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func withLogging(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("-> [%s] %q\n", r.Method, r.URL.String())
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		fmt.Printf("<- [%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}
