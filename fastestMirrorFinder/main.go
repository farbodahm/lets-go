package main

import (
	"fmt"

	"github.com/farbodahm/lets-go/fastestMirrorFinder/fastest_mirror"
	"github.com/farbodahm/lets-go/fastestMirrorFinder/mirrors"
)

func main() {
	mirrorsList := mirrors.GetMirrorsList()
	fastest := fastest_mirror.GetFastestServer(mirrorsList)
	fmt.Println("Fastest:", fastest.Url)
}
