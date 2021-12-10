package main

import "github.com/ptkm1/golang-apiweb.git/server"

func main() {
	server := server.NewServer()

	server.Run()
}
