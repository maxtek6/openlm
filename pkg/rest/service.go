package rest

import (
	"github.com/gorilla/mux"
	"github.com/maxtek6/endpoints-go"
)

type Service struct {
	router        *mux.Router
	usersEndpoint *endpoints.Endpoint
}

func NewService() *Service {
	service := &Service{
		router:        mux.NewRouter(),
		usersEndpoint: endpoints.New(),
	}
	return service
}
