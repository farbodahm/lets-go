package main

import (
	"fmt"

	"github.com/farbodahm/lets-go/fastestMirrorFinder/cmd/fastest_mirror"
)

func main() {
	fastest := fastest_mirror.GetFastestServer()
	fmt.Println("Fastest:", fastest.Url)
}
