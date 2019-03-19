package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"

	"github.com/ardanlabs/service-training/13-context/internal/platform/web"
)

// API constructs an http.Handler with all application routes defined.
func API(db *sqlx.DB, log *log.Logger) http.Handler {

	r := chi.NewRouter()

	p := Products{db: db, log: log}

	r.Post("/v1/products", web.Run(p.Create))
	r.Get("/v1/products", web.Run(p.List))
	r.Get("/v1/products/{id}", web.Run(p.Get))

	return r
}