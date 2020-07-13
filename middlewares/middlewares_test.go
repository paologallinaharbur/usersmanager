package middlewares

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMiddlewareUI(t *testing.T) {

	test := UIMiddleware(nil)
	req, _ := http.NewRequest("GET", "http://test:90/", nil)
	response := httptest.NewRecorder()
	test.ServeHTTP(response, req)

	//We expect the handler to redirect us
	assert.Equal(t, 302, response.Code)
	assert.Equal(t, "/swagger-ui/", response.HeaderMap["Location"][0])

}

func TestMiddlewarePrometheus(t *testing.T) {

	test := PrometheusMiddleware(nil)
	req, _ := http.NewRequest("GET", "http://test:90/metrics", nil)
	response := httptest.NewRecorder()
	test.ServeHTTP(response, req)

	body, _ := ioutil.ReadAll(response.Body)
	assert.True(t, strings.Contains(string(body), "HELP go_gc_duration_seconds"))

}
