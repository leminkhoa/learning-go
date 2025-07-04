package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", RequestIPLogging(TimeHandler()))
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

// Middleware: logs IP address for each request
func RequestIPLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIP := getIP(r)
		logger := createJsonLogging()
		logger.Info("IP Address", slog.String("ip", userIP))
		next.ServeHTTP(w, r)
	})
}

// Handler: responds with time in JSON or text
func TimeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dayOfWeek, dayOfMonth, month, year, hour, minute, second := getCurrentTime()
		acceptHeader := r.Header.Get("Accept")
		if strings.Contains(strings.ToLower(acceptHeader), "application/json") {
			w.Header().Set("Content-Type", "application/json")
			resp := map[string]interface{}{
				"day_of_week":  dayOfWeek,
				"day_of_month": dayOfMonth,
				"month":        month,
				"year":         year,
				"hour":         hour,
				"minute":       minute,
				"second":       second,
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
		// Default: plain text
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s, %d %s %d %02d:%02d:%02d\n", dayOfWeek, dayOfMonth, month, year, hour, minute, second)
	})
}

func createJsonLogging() *slog.Logger {
	options := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stderr, options)
	return slog.New(handler)
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		parts := strings.Split(forwarded, ",")
		return strings.TrimSpace(parts[0])
	}
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func getCurrentTime() (string, int, string, int, int, int, int) {
	currentTime := time.Now().UTC()
	dayOfWeek := currentTime.Weekday().String()
	dayOfMonth := currentTime.Day()
	month := currentTime.Month().String()
	year := currentTime.Year()
	hour := currentTime.Hour()
	minute := currentTime.Minute()
	second := currentTime.Second()
	return dayOfWeek, dayOfMonth, month, year, hour, minute, second
}
