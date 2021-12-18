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

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/items/:itemId", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")                          // get id
		itemId := p.ByName("itemId")                  // get itemId
		output := "Product " + id + " Item " + itemId // output format
		fmt.Fprint(rw, output)
	})

	// create testing (access localhost:3000/product/1/items/1)
	request := httptest.NewRequest("GET", "http://localhost:3000/product/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// assert.Equal(t, ExpectedResult, Result)
	assert.Equal(t, "Product 1 Item 1", string(body))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")   // get id
		output := "Image : " + image // output format
		fmt.Fprint(rw, output)
	})

	// create testing (access localhost:3000/images/profile/xenosty.png)
	request := httptest.NewRequest("GET", "http://localhost:3000/images/profile/xenosty.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// assert.Equal(t, ExpectedResult, Result)
	assert.Equal(t, "Image : /profile/xenosty.png", string(body))
}
