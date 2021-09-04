// Debian fastest mirror finder
//
// The purpose of this application is to:
//
// 1) First scrap all mirrors from Debian site
//
// 2) Find the fastest among them
//
// 3) Present the answer using REST api
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: 0.0.1
//     License: GPL V3 https://github.com/farbodahm/lets-go/blob/main/LICENSE
//     Contact: Farbod Ahmadian<farbodahmadian2014@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
// swagger:meta
package main

import (
	"log"

	"github.com/farbodahm/lets-go/fastestMirrorFinder/internal/api"
)

func main() {
	log.Fatal(api.RunServer("8080"))
}
