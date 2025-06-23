package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gitgud5/shiny-octo-train/internal/app"
)

func main() {

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("We are running our appliction")

	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	http.HandleFunc("/health", HealthCheck)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal("Something happen while running the server")
	}

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
