package app

import (
	"net/http"

	"github.com/alvaro259818/go-post-api/app/database"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     database.PostDB
}

func New() *App {
	app := &App{
		Router: mux.NewRouter(),
	}
	app.initRoutes()
	return app
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods(http.MethodGet)
	a.Router.HandleFunc("/api/posts", a.CreatePostHandler()).Methods(http.MethodPost)
	a.Router.HandleFunc("/api/posts", a.GetPostHandler()).Methods(http.MethodGet)
}
