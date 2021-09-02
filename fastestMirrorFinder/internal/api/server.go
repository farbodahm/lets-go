package api

import "net/http"

func RunServer(port string) error {
	return http.ListenAndServe(":"+port, NewMuxRouter())
}
