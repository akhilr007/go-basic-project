package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/akhilr007/fem-api-project/internal/app"
	"github.com/akhilr007/fem-api-project/internal/routes"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 8080, "backend server port for go")
	flag.Parse()

	app, err := app.NewApplication();
	if err != nil {
		panic(err)
	}

	r := routes.SetUpRoutes(app)
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: r,
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
