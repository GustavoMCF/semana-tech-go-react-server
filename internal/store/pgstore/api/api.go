package api

import (
	"net/http"

	"github.com/GustavoMCF/semana-tech-go-react-server/internal/store/pgstore"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/rooms", func(r chi.Router) {
			r.Post("/", a.handleCreateRoom)
			r.Get("/", a.handleGetRooms)

			r.Route("/{room_id}/messages", func(r chi.Router) {
				r.Post("/", a.handleCreateMessage)
				r.Get("/", a.handleGetRoomMessages)
			})
		})
	})

	a.r = r
	return a

}

func (h apiHandler) handleCreateRoom(w http.ResponseWriter, r *http.Request)      {}
func (h apiHandler) handleGetRooms(w http.ResponseWriter, r *http.Request)        {}
func (h apiHandler) handleCreateMessage(w http.ResponseWriter, r *http.Request)   {}
func (h apiHandler) handleGetRoomMessages(w http.ResponseWriter, r *http.Request) {}
