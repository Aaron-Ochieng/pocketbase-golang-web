package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

func timeParser(dateTimeStr string) (parsedTime time.Time) {
	time_layout := "2006-01-02 15:04:05.000Z"
	parsedTime, err := time.Parse(time_layout, dateTimeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	return parsedTime
}

func Home(se *core.ServeEvent, tmpl *template.Registry) {

	se.Router.GET("/", func(e *core.RequestEvent) error {
		post_collection, err := e.App.FindAllRecords("posts")
		if err != nil {
			return err
		}

		posts := []PostModel{}

		for _, val := range post_collection {
			post := PostModel{}

			post.Images = val.GetStringSlice("image")
			post.Content = val.GetString("content")
			post.Id = val.Id
			created_at := fmt.Sprintf("%s", val.Collection().Created)
			post.CreatedAt = timeParser(created_at)
			updated_at := fmt.Sprintf("%s", val.Collection().Updated)
			post.UpdatedAt = timeParser(updated_at)

			posts = append(posts, post)
		}

		fmt.Println(posts)
		html, err := tmpl.LoadFiles(
			"templates/layout.html",
			"templates/home.html",
		).Render(map[string]any{
			"title": "Home",
			"posts": posts,
		})

		if err != nil {
			return e.NotFoundError("", err)
		}

		return e.HTML(http.StatusOK, html)
	})
}
