package main

import (
	"github.com/gitgud5/shiny-octo-train/internal/app"
)

func main() {

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("We are running our appliction")

}
