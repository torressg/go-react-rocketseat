package api

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	"github.com/torressg/go-react-rocketseat/internal/store/pgstore/pgstore"
)

type apiHandler struct {
	q        *pgstore.Queries
	r        *chi.Mux
	upgrader websocket.Upgrader
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
		q:        q,
		upgrader: websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/subscribe/{room_id}", a.handleSubscribe)

	r.Route("/api", func(r chi.Router) {
		r.Route("/rooms", func(r chi.Router) {
			r.Post("/", a.handleCreateRoom)
			r.Get("/", a.handleGetRooms)

			r.Route("/{room_id}/messages", func(r chi.Router) {
				r.Post("/", a.handleCreateRoomMessage)
				r.Get("/", a.handleGetRoomsMessages)

				r.Route("/{message_id}", func(r chi.Router) {
					r.Get("/", a.handleGetRoomMessage)
					r.Patch("/react", a.handleReactionToMessage)
					r.Delete("/react", a.handleRemoveReactionFromMessage)
					r.Patch("/answer", a.handleMarkMessageAsAnswered)
				})
			})
		})
	})

	a.r = r
	return a
}

func (h apiHandler) handleSubscribe(w http.ResponseWriter, r *http.Request) {

}
func (h apiHandler) handleCreateRoom(w http.ResponseWriter, r *http.Request)                {}
func (h apiHandler) handleGetRooms(w http.ResponseWriter, r *http.Request)                  {}
func (h apiHandler) handleGetRoomsMessages(w http.ResponseWriter, r *http.Request)          {}
func (h apiHandler) handleCreateRoomMessage(w http.ResponseWriter, r *http.Request)         {}
func (h apiHandler) handleGetRoomMessage(w http.ResponseWriter, r *http.Request)            {}
func (h apiHandler) handleReactionToMessage(w http.ResponseWriter, r *http.Request)         {}
func (h apiHandler) handleRemoveReactionFromMessage(w http.ResponseWriter, r *http.Request) {}
func (h apiHandler) handleMarkMessageAsAnswered(w http.ResponseWriter, r *http.Request)     {}
