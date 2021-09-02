package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/farbodahm/lets-go/fastestMirrorFinder/pkg/fastest_mirror"
	"github.com/farbodahm/lets-go/fastestMirrorFinder/pkg/mirrors"
)

func GetFastestMirror(w http.ResponseWriter, r *http.Request) {
	fastestMirror := fastest_mirror.GetFastestServer(mirrors.GetMirrorsList())
	responseJson, err := json.Marshal(fastestMirror)
	if err != nil {
		log.Fatalln("Error: Couldn't serialize fastest mirror:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(responseJson); err != nil {
		log.Fatal("Error: Couldn't write to the request", err)
	}
}
