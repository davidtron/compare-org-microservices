package main

import "testing"

import (
  "fmt"
  "io"
  "net/http/httptest"
)

func TestGetAllEvents(t *testing.T) {
  // (w http.ResponseWriter, r *http.Request)
  // w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(events)

  // var dummyEvent = event{1, "foobar", "baz"}
  // events = append(events, dummyEvent)


  req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	GetAllEvents(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
  
}
