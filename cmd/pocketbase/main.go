package main

import (
	_ "github.com/burneyhoel/pocketbase/migrations"
	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		panic(err)
	}
}
