package Golang_Web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleWare struct {
	Handler http.Handler
}

func (middleware *LogMiddleWare) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Test Executed")
		fmt.Fprint(writer, "Hello Test Middleware")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("panic Executed")
		panic("Error")
	})

	logMiddleWare := &LogMiddleWare{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleWare,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
