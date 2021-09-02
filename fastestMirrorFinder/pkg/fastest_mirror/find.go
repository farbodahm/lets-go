package fastest_mirror

import (
	"net/http"
	"time"
)

const MAX_TIME_OUT_MILISECONDS = 2000

type FastestMirror struct {
	Url     string        `json:"url"`
	Latency time.Duration `json:"latency"`
}

func GetFastestServer(mirrorsList []string) FastestMirror {
	timeoute := time.Duration(MAX_TIME_OUT_MILISECONDS) * time.Millisecond
	client := http.Client{
		Timeout: timeoute,
	}
	result := FastestMirror{
		Latency: timeoute,
	}

	for _, url := range mirrorsList {
		current := time.Now()
		_, err := client.Get(url)
		latency := time.Since(current)

		if err == nil && latency < result.Latency {
			result = FastestMirror{url, latency}
		}
	}

	return result
}
