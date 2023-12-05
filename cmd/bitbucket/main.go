package main

import (
	"log"
	"net/http"

	"github.com/zomgra/bitbucket/internal/routes"
)

func main() {
	r := routes.CreateRoute()
	log.Fatal(http.ListenAndServe(":8000", r))
}
