package models

import (
	"encoding/json"
	"net/http"
	"time"
)

type Blogs struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type APIError struct {
	ApiError string `json:"error"`
}

func (req *Blogs) CreateBlog(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	//validations
	if req.Title != "Oluwatosin" {
		err := &APIError{ApiError: "Sorry the tile is wrong"}
		json.NewEncoder(w).Encode(err)
		return
	}

	// store into the database

	// return responses
	res := &Blogs{
		Title: req.Title,
		Body:  req.Body,
	}
	json.NewEncoder(w).Encode(res)
}
