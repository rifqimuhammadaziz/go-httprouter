package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	// create panic handler
	router.PanicHandler = func(rw http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(rw, "Panic : ", error) // output in html
	}

	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Under Maintenance")
	})

	// create testing
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// assert.Equal(t, ExpectedResult, Result)
	assert.Equal(t, "Panic : Under Maintenance", string(body))
}
