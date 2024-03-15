package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	ctx context.Context
	db  *sqlx.DB
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	db, err := sqlx.Connect("mysql", "root:1234@tcp(localhost:3306)/regio")
	if err != nil {
		panic(err)
	}
	a.ctx = ctx
	a.db = db
}

type Greeting struct {
	Message string
}

func (a *App) Greet(name string) *Greeting {
	return &Greeting{
		Message: fmt.Sprintf("Hello, %s!", name),
	}
}
