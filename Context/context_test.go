package main

import (
	// "context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	// data := "hello, world"
	// svr := Server(&SpyStore{data})

	// request := httptest.NewRequest(http.MethodGet, "/", nil)
	// response := httptest.NewRecorder()

	// svr.ServeHTTP(response, request)

	// if response.Body.String() != data {
	// 	t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	// }

	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)
	
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
	
		svr.ServeHTTP(response, request)
	
		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	
		if store.cancelled {
			t.Error("it should not have cancelled the store")
		}
	})
}