package api

import (
	"github.com/go-chi/chi"
)

// InitAPIRoutes sets project API routes
func InitAPIRoutes(r chi.Router) {
	/*
		|--------------------------------------------------------------------------
		| Users Routes
		|--------------------------------------------------------------------------
		|
	*/
	r.Route("/pertanyaan", func(r chi.Router) {
		r.Get("/", GetJawaban)
		r.Post("/", StoreQNA)
		r.Get("/hewo", HelloWorld)
	})

	r.Route("/pengguna", func(r chi.Router) {
		r.Get("/", GetPengguna)
		r.Post("/", StorePengguna)
	})

	r.Route("/keluhan", func(r chi.Router) {
		r.Post("/", StoreKeluhan)
	})

	r.Route("/employee", func(r chi.Router) {
		r.Get("/", GetAllEmployee)
	})
}
