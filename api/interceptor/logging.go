package interceptor

import (
	"log"
	"net/http"
)

func LoggingInterCeptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.RequestURI, r.Method)
	})
}
