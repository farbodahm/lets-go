package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/farbodahm/lets-go/fastestMirrorFinder/internal/api"
	"github.com/farbodahm/lets-go/fastestMirrorFinder/pkg/fastest_mirror"
)

func TestGetFastestMirror(t *testing.T) {
	// Find route's url, method and handler
	var expectedRoute api.Route
	for _, route := range api.Routes {
		if route.Name == "GetFastestMirror" {
			expectedRoute = route
			break
		}
	}

	// Create response
	req, err := http.NewRequest(expectedRoute.Method, expectedRoute.Pattern, nil)
	if err != nil {
		t.Fatal("Error: Couldn't send request in test:", err)
	}
	responseRecorder := httptest.NewRecorder()
	expectedRoute.HandlerFunc(responseRecorder, req)

	// Check the status code
	if responseRecorder.Code != http.StatusOK {
		t.Error("Handler returned wrong status code: got",
			responseRecorder.Code, "want", http.StatusOK,
		)
	}

	// Check the response body
	var returnedBody fastest_mirror.FastestMirror
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &returnedBody)
	if err != nil {
		t.Error("Couldn't unmarshal", responseRecorder.Body.String())
	}

	if returnedBody.Url == "" {
		t.Error("Returned Url is not. Got:", returnedBody.Url)
	}
	if returnedBody.Latency <= time.Duration(0) ||
		returnedBody.Latency >= time.Duration(fastest_mirror.MAX_TIME_OUT_MILISECONDS)*time.Millisecond {
		t.Error(
			"Latency should be between 0 and",
			fastest_mirror.MAX_TIME_OUT_MILISECONDS, "miliseconds",
			"got", returnedBody.Latency,
		)
	}
}
