package config

import (
	"GoAppOne/app/api"
	"encoding/json"
	"net/http"
)

func (server *Server) initializeRoutes() {
	/*
	 Route Home
	*/

	server.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"Message": "Hey, welcome to my REST API",
		}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")
	server.Router.HandleFunc("/user", api.UserGet).Methods("GET")
	server.Router.HandleFunc("/user", api.UserPost).Methods("POST")
	server.Router.HandleFunc("/user/{id:[0-9]+}", api.UserGet).Methods("GET")
	server.Router.HandleFunc("/user/{id:[0-9]+}", api.UserPut).Methods(http.MethodPut)
	server.Router.HandleFunc("/user/{id:[0-9]+}", api.UserDelete).Methods("DELETE")

	server.Router.HandleFunc("/book", api.BookGet).Methods("GET")
	server.Router.HandleFunc("/book", api.BookPost).Methods("POST")
	server.Router.HandleFunc("/book/{id:[0-9]+}", api.BookGet).Methods("GET")
	server.Router.HandleFunc("/book/{id:[0-9]+}", api.BookPut).Methods(http.MethodPut)
	server.Router.HandleFunc("/book/{id:[0-9]+}", api.BookDelete).Methods("DELETE")
}
