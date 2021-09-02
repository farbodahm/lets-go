package mirrors

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const (
	MIRRORS_LIST_URL             = "https://www.debian.org/mirror/list"
	DEBIAN_MIRRORS_REGEX_PATTERN = "http://.*?debian.org/debian"
	MAX_MIRROR_NUMBER_TO_GET     = 39 // Max = 39
)

func GetMirrorsList() []string {
	response, err := http.Get(MIRRORS_LIST_URL)
	if err != nil {
		log.Fatalln("Error: Could not get site: ", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error: Could not read body: ", err)
	}

	regexComp, err := regexp.Compile(DEBIAN_MIRRORS_REGEX_PATTERN)
	if err != nil {
		log.Fatalln("Error: Could not compile regex pattern: ", err)
	}

	return regexComp.FindAllString(string(body), MAX_MIRROR_NUMBER_TO_GET)
}
