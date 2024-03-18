package main

import (
	"embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello World", string(body))
}

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		text := "Product " + id
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
}

func TestParamsMulti(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId",
		func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			id := params.ByName("id")
			itemId := params.ByName("itemId")
			text := "Product " + id + " Item " + itemId
			fmt.Fprint(writer, text)
		})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 1", string(body))
}

func TestParamsAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image",
		func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			image := params.ByName("image")
			text := "Image : " + image
			fmt.Fprint(writer, text)
		})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image : /small/profile.png", string(body))
}

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello HTTP Router", string(body))
}

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
		fmt.Fprint(writer, "Panic : ", err)
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Something went wrong")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic : Something went wrong", string(body))
}

func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Not Found")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Not Found", string(body))
}

func TestNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Method Not Allowed")
	})

	router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Method Not Allowed", string(body))
}

// Middleware HTTP Router

type LogMiddleware struct {
	handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Log Middleware")
	middleware.handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Middleware")
	})

	middleware := LogMiddleware{router}

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(body))
}
