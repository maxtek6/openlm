// Copyright (c) 2024 Maxtek Consulting
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
