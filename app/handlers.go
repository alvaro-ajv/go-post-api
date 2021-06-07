package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alvaro259818/go-post-api/app/models.go"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Welcome to Post API")
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := parse(rw, r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(rw, r, nil, http.StatusBadRequest)
			return
		}

		// Create the post
		p := &models.Post{
			ID:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		// Save in DB
		err = a.DB.CreatePost(p)
		if err != nil {
			log.Printf("Cannot save the post in DB. err=%v\n", err)
			sendResponse(rw, r, nil, http.StatusInternalServerError)
			return
		}
		resp := mapPostToJSON(p)
		sendResponse(rw, r, resp, http.StatusOK)
	}
}

func (a *App) GetPostHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		posts, err := a.DB.GetPosts()
		if err != nil {
			log.Printf("Cannot get posts, err=%v", err)
			sendResponse(rw, r, nil, http.StatusInternalServerError)
			return
		}
		var resp = make([]models.JsonPost, len(posts))
		for index, post := range posts {
			resp[index] = mapPostToJSON(post)
		}

		sendResponse(rw, r, resp, http.StatusOK)
	}
}
