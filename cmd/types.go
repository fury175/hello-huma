package main

type application struct {
	config config
}

type config struct {
	ADDR string `env:"ADDR,required" envDefault:":3001"`
	ENV  string `env:"ENV" envDefault:"development"`
}
