package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that a GET request to / endpoint returns the home page and a 200 OK
func TestShowIndexPage(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

// Test that a GET request to /rand endpoint returns an exercise and a 200 OK
func TestRandomExercise(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/rand", getRandomExercise)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/rand", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title has a URL as a ref in the title
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<a href=") > 0

		return statusOK && pageOK
	})
}

// Test that a GET request to a random exercise returns in JSON
// format when the Accept header is set to application/json
func TestArticleListJSON(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/rand", getRandomExercise)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/rand", nil)
	req.Header.Add("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the response is JSON which can be converted to a struct
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var e exercise
		err = json.Unmarshal(p, &e)

		return err == nil && e.Name != "" && statusOK
	})
}

// Test that a GET request to a random exercise returns in XML
// format when the Accept header is set to application/xml
func TestArticleXML(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/rand", getRandomExercise)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/rand", nil)
	req.Header.Add("Accept", "application/xml")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the response is XML which can be converted to a struct
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var e exercise
		err = xml.Unmarshal(p, &e)

		return err == nil && e.Name != "" && statusOK
	})
}
