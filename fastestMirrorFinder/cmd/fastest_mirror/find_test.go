package fastest_mirror_test

import (
	"testing"
	"time"

	"github.com/farbodahm/lets-go/fastestMirrorFinder/cmd/fastest_mirror"
	"github.com/farbodahm/lets-go/fastestMirrorFinder/mirrors"
)

func TestFindingFastestMirror(t *testing.T) {
	mirrorsList := mirrors.GetMirrorsList()
	result := fastest_mirror.GetFastestServer(mirrorsList)

	if result.Latency <= time.Duration(0) ||
		result.Latency >= time.Duration(fastest_mirror.MAX_TIME_OUT_MILISECONDS)*time.Millisecond {
		t.Error(
			"Latency should be between 0 and",
			fastest_mirror.MAX_TIME_OUT_MILISECONDS, "miliseconds",
			"got", result.Latency,
		)
	}

	if result.Url == "" {
		t.Error(
			"Url can not be empity!",
		)
	}
}
