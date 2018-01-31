package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route func(s *Server)

func AddRoutes(s *Server) {
	s.addRoute("HealthCheck", "GET", "/dbstatus", s.DbStatus)
	s.addRoute("GetPricesData", "GET", "/prices/from/{fromsymbol}/fromts/{fromTime}", s.ExtractData)
}

func (s *Server) addRoute(name string, method string, pattern string, fn http.HandlerFunc) {
	var handler http.Handler
	handler = s.WrapRequest(fn, name)

	s.Server.Handler.(*mux.Router).Methods(method).Path(pattern).Name(name).Handler(handler)

	if pattern != "/" {
		// Add support for OPTIONS requests
		fn = s.Options
		handler = s.WrapRequest(fn, name)

		s.Server.Handler.(*mux.Router).Methods("OPTIONS").Path(pattern).Name(name).Handler(handler)
	}
}
