package server

import (
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/web/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router     *chi.Mux
	server     *http.Server
	accountSvc *service.AccountService
	port       string
}

func NewServer(svc *service.AccountService, port string) *Server {
	return &Server{
		router:     chi.NewRouter(),
		accountSvc: svc,
		port:       port,
	}
}

func (s *Server) SetRoutes() {
	actHandler := handlers.NewAccountHandler(s.accountSvc)

	s.router.Post("/accounts", actHandler.Create)
	s.router.Get("/accounts", actHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
