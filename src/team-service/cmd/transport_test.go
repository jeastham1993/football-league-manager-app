package main

import (
	"net/http"
	"testing"
)

func TestCanMakeHandler(t *testing.T) {
	var h http.Handler
	{
		endpoints := MakeEndpoints(NewMockService())
		h = MakeHandler(endpoints)
	}

	if h == nil {
		t.Fatalf("Handler not properly initialised")
	}
}
