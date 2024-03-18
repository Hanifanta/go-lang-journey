package Golang_Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello World!")
	} else {
		fmt.Fprintf(writer, "Hello %s!", name)
	}
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s!", firstName, lastName)
}

func MultipleValuesParameter(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Hanif&name=Al&name=Irsyad", nil)
	recorder := httptest.NewRecorder()

	MultipleValuesParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
