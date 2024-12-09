package main

import (
	"log"
	"pocketbase-web/cmd/routes"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		registry := template.NewRegistry()

		routes.Home(se, registry)
		routes.Register(se, registry)
		routes.CreatePost(se, registry)

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
