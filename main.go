package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/akhilr007/fem-api-project/internal/app"
)

func main() {

	app, err := app.NewApplication();
	if err != nil {
		panic(err)
	}

	app.Logger.Println("We have started our server")

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr: ":8080",
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}