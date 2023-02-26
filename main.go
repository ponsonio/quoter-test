package main

import (
	"gitlab.com/ponsonio/quoter-test/mortgage/server"
	"log"
	"net/http"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
