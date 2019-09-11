package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

//GetRouter fetches the chi router
func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
	)

	router.Route("/v1", func(router chi.Router) {
		router.Mount("/api/mongo_app", httpRoutes())
	})

	return router
}

func httpRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{userID}", GetProfile)
	router.Delete("/{userID}", DeleteProfile)
	router.Put("/{userID}", UpdateProfile)
	router.Post("/", CreateProfile)
	router.Get("/", GetAllProfiles)

	return router
}
