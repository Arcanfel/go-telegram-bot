package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

const (
	serverPort = ":9000"
)

// Server is a simple struct that handles server related
// things like routing
type Server struct {

	Port string
	router *mux.Router
}

// NewServer returns newly confgured server with default const
// values declared in const section in the beginning of the file
func NewServer() *Server {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", RootHandler)

	s := &Server{Port: serverPort, router: r}

	return s
}

// RootHandler does things
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello friend\n"))
}

// Start start the server
func (s *Server) Start() {
	// Bind to a port and pass our router in

	fmt.Printf("Started server at %s", s.Port)
	http.ListenAndServe(s.Port, s.router)
}