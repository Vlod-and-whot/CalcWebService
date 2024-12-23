package main

import (
	"CalculationWebService/Packages/Server"
)

func main() {
	address := "localhost:8080"
	server := Server.NewServer(address)
	server.Run()
}
