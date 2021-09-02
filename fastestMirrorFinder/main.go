package main

import (
	"log"

	"github.com/farbodahm/lets-go/fastestMirrorFinder/internal/api"
)

func main() {
	log.Fatal(api.RunServer("8080"))
}
