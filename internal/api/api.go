package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/torressg/go-react-rocketseat/internal/store/pgstore/pgstore"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

// ServeHTTP implements http.Handler.
func (h apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

func (h apiHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
	}
	r := chi.NewRouter()

	a.r = r
	return a
}
