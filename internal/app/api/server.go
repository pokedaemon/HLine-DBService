package api

import (
	"hlservice-db/internal/app/storage"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router  *mux.Router
	logger  *logrus.Logger
	storage storage.Storage
}

func New(store storage.Storage) *server {
	s := &server{
		router:  mux.NewRouter(),
		logger:  logrus.New(),
		storage: store,
	}

	s.setRouter()

	return s
}

func (s *server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) setRouter() {

}
