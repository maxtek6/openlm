package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxtek6/endpoints-go"
)

// Service is an HTTP router for all REST endpoints
type Service struct {
	router         *mux.Router
	usersEndpoint  *endpoints.Endpoint
	userEndpoint   *endpoints.Endpoint
	statusEndpoint *endpoints.Endpoint
}

func NewService() *Service {
	service := &Service{
		router:         mux.NewRouter(),
		usersEndpoint:  endpoints.New(),
		userEndpoint:   endpoints.New(),
		statusEndpoint: endpoints.New(),
	}

	_ = service.router.Handle("/v1/users", service.usersEndpoint)
	_ = service.router.Handle("/v1/users/{userid}", service.userEndpoint)
	_ = service.router.Handle("/v1/products", service.userEndpoint)
	_ = service.router.Handle("/v1/product/{productid}", service.userEndpoint)
	_ = service.router.Handle("/v1/users/{userid}", service.userEndpoint)
	_ = service.router.Handle("/v1/status", service.statusEndpoint)
	return service
}

func (s *Service) HandleUsers(onPost http.HandlerFunc) error {
	handlers := map[string]http.HandlerFunc{
		http.MethodPost: onPost,
	}
	err := setupEndpoint(s.usersEndpoint, handlers)
	if err != nil {
		return fmt.Errorf("service.HandleUsers: %v", err)
	}
	return nil
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func setupEndpoint(endpoint *endpoints.Endpoint, handlers map[string]http.HandlerFunc) error {
	if endpoint == nil {
		return errors.New("nil endpoint")
	}
	for method, handler := range handlers {
		err := endpoint.AddMethod(method, handler)
		if err != nil {
			return fmt.Errorf("invalid %s handler: %v", method, err)
		}
	}
	return nil
}
