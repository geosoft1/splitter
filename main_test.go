package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	c.Routes = []string{"http://192.168.88.1:8080", "http://192.168.88.2:8080"}
	t.Log(c.Routes)
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()
	_, err := http.Get(ts.URL + "/update?key=0&a=1&b=2")
	if err != nil {
		t.Fatal(err)
	}
}
