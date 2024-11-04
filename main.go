package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type CurrentTime struct {
	Time int64 `json:"time"`
}

func currentTime(now func() time.Time) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timeC := now().Add(1 * time.Second).UnixMilli()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		err := json.NewEncoder(w).Encode(CurrentTime{Time: timeC})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			os.Exit(1)
		}
	})
	mux.HandleFunc("GET /current_time", currentTime(time.Now))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
