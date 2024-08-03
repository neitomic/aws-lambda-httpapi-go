package main

import (
	"io"
	"lambda_app/pkg/lambdaapi"
	"net/http"
)

func main() {
	app := lambdaapi.NewHttpApiApp()
	app.HandlerFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Here is what I got: "))
		w.Write(data)
	})
	app.Start()
}
