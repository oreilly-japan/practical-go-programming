package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// server-test
func TestServerRequest(t *testing.T) {
	s := httptest.NewServer(initServer())
	r, err := http.Get(s.URL + "/fortune")
	if err != nil {
		t.Errorf("http get err should be nil: %v", err)
		return
	}
	defer r.Body.Close()
	var j map[string]string
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		t.Errorf("json decode err should be nil: %v", err)
		return
	}
	if j["fortune"] != "大吉" {
		t.Errorf("result should be 大吉, but %s", j["fortune"])
	}
}

// server-test

// handler-test
func TestHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/fortune", nil)
	w := httptest.NewRecorder()
	fortuneHandler(w, r)
	var j map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &j); err != nil {
		t.Errorf("json unmarshal err should be nil: %v", err)
		return
	}
	if j["fortune"] != "大吉" {
		t.Errorf("result should be 大吉, but %s", j["fortune"])
	}
}

// handler-test
