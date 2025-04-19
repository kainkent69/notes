package middlewares

import (
	"log"
	"net/http"
	"time"
)

type wrapWritter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrapWritter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(w.statusCode)
	w.statusCode = status
}

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &wrapWritter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		log.Println(wrapper.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
