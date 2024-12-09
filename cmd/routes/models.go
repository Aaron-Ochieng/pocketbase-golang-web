package routes

import "time"

type PostModel struct {
	Id        string
	Content   string
	Images    []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
