package web

import (
	"server/web/handlers"

	"github.com/go-chi/chi/v5"
)

func Web(r *chi.Mux) {
	h := handlers.Handlers{}

	r.Get("/", h.Welcome)
	r.Get("/abc/{id}", h.Abc)
	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
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
