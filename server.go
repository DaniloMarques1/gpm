package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server interface {
	Run() error
}

type server struct {
	router           chi.Router
	masterController MasterController
}

func NewServer() (Server, error) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	masterController, err := NewMasterControler()
	if err != nil {
		return nil, err
	}

	return &server{router, masterController}, nil
}

func (s *server) Run() error {
	s.router.Get("/", s.HelloWorld)
	s.router.Post("/master/signup", s.masterController.SignUp)
	s.router.Post("/master/signin", s.masterController.SignIn)

	log.Printf("Starting server\n")
	return http.ListenAndServe(":3000", s.router) // TODO: move port somewhere else
}

func (s *server) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world\n"))
}
