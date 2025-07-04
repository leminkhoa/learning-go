package main

/**
Write a small web server that returns the current time in RFC 3339 format when
you send it a GET command. You can use a third-party module if you’d like.
**/

import (
	"net/http"
	"time"
)

type RFC3339Handler struct{}

func (h RFC3339Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	currentTime := time.Now().UTC().Format(time.RFC3339)

	w.Write([]byte("current time is " + currentTime))
}

func main() {
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      RFC3339Handler{},
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

}
