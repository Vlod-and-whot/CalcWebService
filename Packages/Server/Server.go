package Server

import (
	"CalculationWebService/Packages/Handler"
	"net/http"
)

type Server struct {
	Address string
}

func NewServer(address string) *Server {
	return &Server{
		Address: address,
	}
}

func (s *Server) Run() error {
	http.HandleFunc("/api/v1/calculate", Handler.CalcHandler)
	return http.ListenAndServe(s.Address, nil)
}
