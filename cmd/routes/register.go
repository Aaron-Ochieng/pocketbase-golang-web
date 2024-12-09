package routes

import (
	"fmt"
	"net/http"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

func Register(se *core.ServeEvent, tmpl *template.Registry) {

	se.Router.GET("/register", func(e *core.RequestEvent) error {
		html, err := tmpl.LoadFiles(
			"templates/layout.html",
			"templates/register.html",
		).Render(nil)

		if err != nil {
			return e.NotFoundError("", err)
		}

		return e.HTML(http.StatusOK, html)
	})

	se.Router.POST("/register", func(e *core.RequestEvent) error {
		username := e.Request.FormValue("username")
		email := e.Request.FormValue("email")
		password := e.Request.FormValue("password")
		confirm := e.Request.FormValue("confirm")

		// html, err := tmpl.LoadFiles(
		// 	"templates/layout.html",
		// 	"templates/register.html",
		// ).Render(nil)
		//
		fmt.Println("%s %s %s %s\n", username, email, password, confirm)

		return e.Redirect(http.StatusFound, "/")
	})
}
