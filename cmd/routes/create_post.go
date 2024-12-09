package routes

import (
	"fmt"
	"net/http"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

func CreatePost(se *core.ServeEvent, tmpl *template.Registry) {

	se.Router.GET("/new_post", func(e *core.RequestEvent) error {
		html, err := tmpl.LoadFiles(
			"templates/layout.html",
			"templates/new_post.html",
		).Render(map[string]any{
			"title": "New Post",
		})

		if err != nil {
			return e.NotFoundError("", err)
		}

		return e.HTML(http.StatusOK, html)
	})

	se.Router.POST("/new_post", func(e *core.RequestEvent) error {
		content := e.Request.FormValue("content")
		fmt.Println(content)

		post_collection, err := e.App.FindCollectionByNameOrId("posts")
		if err != nil {
			return err
		}

		record := core.NewRecord(post_collection)
		record.Set("content", content)

		if err = e.App.Save(record); err != nil {
			return err
		}
		return e.Redirect(http.StatusFound, "/")
	})
}
