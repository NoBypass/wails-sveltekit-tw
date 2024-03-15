package main

import (
	"context"
	"fmt"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type Greeting struct {
	Message string
}

func (a *App) Greet(name string) *Greeting {
	return &Greeting{
		Message: fmt.Sprintf("Hello, %s!", name),
	}
}
