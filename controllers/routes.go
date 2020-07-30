package controllers

import "github.com/dennypradipta/cariliquid/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Root)).Methods("GET")
	
	// User Routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{userID}",  middlewares.SetMiddlewareJSON(s.GetUserByID)).Methods("GET")
	s.Router.HandleFunc("/users/{userID}",  middlewares.SetMiddlewareJSON(s.UpdateUserByID)).Methods("PATCH")
	s.Router.HandleFunc("/users/{userID}",  middlewares.SetMiddlewareJSON(s.DeleteUserByID)).Methods("DELETE")
}
