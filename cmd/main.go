package main

import (
	"log"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

const version = "0.0.1"

func main() {
	cfg, err := env.ParseAs[config]()
	if err != nil {
		log.Fatal(err)
	}

	app := application{config: cfg}

	if err := app.run(app.mount()); err != nil {
		log.Fatal(err)
	}
}
