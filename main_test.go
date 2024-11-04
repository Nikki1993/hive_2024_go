package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func now() time.Time {
	return time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC)
}

func TestCurrentTime(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/current_time", nil)
	recorder := httptest.NewRecorder()

	handler := currentTime(now)
	handler(recorder, req)

	result := recorder.Result()
	if result.Header.Get("Cache-Control") != "no-cache" {
		t.Error("Cache control header should not be empty")
	}

	ct := CurrentTime{}
	_ = json.NewDecoder(result.Body).Decode(&ct)

	if ct.Time != now().UnixMilli() {
		t.Error("time is wrong")
	}
}
