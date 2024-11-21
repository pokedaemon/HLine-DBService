package apiserver

import (
	"hlservice-db/internal/app/storage"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

// New create new API service
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start
func (a *API) Start() error {
	if err := a.configureLogger(); err != nil {
		return err
	}
	if err := a.configureStorage(); err != nil {
		return err
	}

	a.configureRouter()

	a.logger.Info("Starting API Service!")
	return http.ListenAndServe(a.config.BindAddr, a.router)
}

// configureLogger parse a log level
func (a *API) configureLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)

	return nil
}

// configureRouter handle all endpoints what we need
func (a *API) configureRouter() {
	a.router.HandleFunc("/hello", a.handleHello())
}

func (a *API) configureStorage() error {
	st := storage.New(a.config.Storage)
	if err := st.Open(); err != nil {
		return err
	}

	a.storage = st

	return nil
}

func (a *API) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info(r.Method, " /hello")
		io.WriteString(w, "Hello from API!")
	}
}
