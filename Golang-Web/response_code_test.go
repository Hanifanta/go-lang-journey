package Golang_Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Name is empty")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hello %s!", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(bodyString)
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=hanif", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(bodyString)
}
