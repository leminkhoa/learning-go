package main

import (
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

/**
Write a small middleware component that uses JSON structured logging to log
the IP address of each incoming request to your web server.
**/

func main() {
	mux := http.NewServeMux()

	jsonLogger := createJsonLogging()

	// to apply the middleware to just the single route
	mux.Handle("/hello", RequestIPLogging(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello!\n"))
		}), jsonLogger),
	)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

}

func RequestIPLogging(h http.Handler, l *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIP := getIP(r)

		h.ServeHTTP(w, r)
		l.Info("IP Address", slog.String("ip", userIP))
	})
}

func createJsonLogging() *slog.Logger {
	options := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stderr, options)
	return slog.New(handler)
}

func getIP(r *http.Request) string {
	// Try to get IP from the X-Forwarded-For header (can contain multiple IPs)
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Sometimes it’s a comma-separated list: client, proxy1, proxy2...
		parts := strings.Split(forwarded, ",")
		return strings.TrimSpace(parts[0])
	}

	// Fallback to X-Real-IP
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // might be just the IP or IP:port
	}
	return ip
}
