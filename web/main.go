package web

import (
	"github.com/go-chi/chi/v5"
	"server/web/handlers"
)

func Web(r *chi.Mux) {
	h := handlers.Handlers{}

	r.Get("/", h.Welcome)
	r.Get("/abc/{id}", h.Abc)
	r.NotFound(h.NotFound)

	// create subroute
	subRouter := chi.NewRouter()
	subRouter.Get("/articles", h.Articles)
	r.Mount("/api", subRouter)

	//routing groups
	r.Group(func(r chi.Router) {
		//you can use middleware here to affect to routes in this group
		r.Get("/post", h.Post)
	})
}
