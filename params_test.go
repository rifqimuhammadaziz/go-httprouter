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

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// get id
		id := p.ByName("id")
		output := "Product " + id // Product 1
		fmt.Fprint(rw, output)
	})

	// create testing (access localhost:3000/product/1)
	request := httptest.NewRequest("GET", "http://localhost:3000/product/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// assert.Equal(t, ExpectedResult, Result)
	assert.Equal(t, "Product 1", string(body))
}
