package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/farbodahm/lets-go/fastestMirrorFinder/internal/api/database"
	"github.com/farbodahm/lets-go/fastestMirrorFinder/pkg/fastest_mirror"
	"github.com/farbodahm/lets-go/fastestMirrorFinder/pkg/mirrors"
)

// swagger:operation GET /fastest-mirror fastestMirror
//
// Get fastest mirror
//
// Find the fastest mirror among Debian mirrors
//
// ---
// produces:
// - application/json
// responses:
//   "200":
//     description: Fastest mirror
//     schema:
//       "$ref": "#/definitions/fastestMirror"
func GetFastestMirror(w http.ResponseWriter, r *http.Request) {
	// Check to see if mirror urls already exist, if not, scrap and save to cache
	isUrlExist, err := database.RedisClient.Exists("mirrors").Result()
	if err != nil {
		log.Fatal("Error: Couldn't execute query to Redis:", err)
	}

	var mirrorsList []string
	if isUrlExist == 1 {
		err = database.RedisClient.LRange("mirrors", 0, -1).ScanSlice(&mirrorsList)
		if err != nil {
			log.Fatal("Error: Couldn't get mirrors:", err)
		}
	} else {
		mirrorsList = mirrors.GetMirrorsList()
		err = database.RedisClient.RPush("mirrors", mirrorsList).Err()
		if err != nil {
			log.Fatal("Error: Couldn't push mirrors:", err)
		}
	}

	// Fetch fastest mirror
	fastestMirror := fastest_mirror.GetFastestServer(mirrorsList)
	responseJson, err := json.Marshal(fastestMirror)
	if err != nil {
		log.Fatalln("Error: Couldn't serialize fastest mirror:", err)
	}

	// Setting Headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(responseJson); err != nil {
		log.Fatal("Error: Couldn't write to the request", err)
	}
}
