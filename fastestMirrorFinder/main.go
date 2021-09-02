package main

import (
	"fmt"

	"github.com/farbodahm/lets-go/fastestMirrorFinder/pkg/fastest_mirror"
	"github.com/farbodahm/lets-go/fastestMirrorFinder/pkg/mirrors"
)

func main() {
	mirrorsList := mirrors.GetMirrorsList()
	fastest := fastest_mirror.GetFastestServer(mirrorsList)
	fmt.Println("Fastest:", fastest.Url)
}
