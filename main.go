package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/akhilr007/fem-api-project/internal/app"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 8080, "backend server port for go")
	flag.Parse()

	app, err := app.NewApplication();
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}

	app.Logger.Printf("Server started running at %d", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}