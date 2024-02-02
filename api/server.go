package api

import "net/http"

type Server struct {
	listenAddr string
}

func NewServer(ListenAddr string) *Server {
	return &Server{
		listenAddr: ListenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/depth", s.depth)
	http.HandleFunc("/order", s.order)
	http.HandleFunc("/balance", s.balance)

	return http.ListenAndServe(s.listenAddr, nil)
}
