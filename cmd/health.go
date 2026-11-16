package main

import "context"

type HealthBody struct {
	Status  string `json:"status"`
	Env     string `json:"environment"`
	Success bool   `json:"success"`
}

type Health struct {
	Body HealthBody
}

func (app *application) healthCheckHandler(ctx context.Context, input *struct{}) (*Health, error) {
	return &Health{Body: HealthBody{Status: "ok", Env: app.config.ENV, Success: true}}, nil
}
