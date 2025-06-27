package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gitgud5/shiny-octo-train/internal/app"
	"github.com/gitgud5/shiny-octo-train/internal/routes"
)

func main() {

	var port int

	flag.IntVar(&port, "port", 8080, "Go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	defer app.DB.Close()

	app.Logger.Println("We are running our appliction on port", port)

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal("Something happen while running the server")
	}

}
