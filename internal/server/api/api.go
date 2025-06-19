package api

import "net/http"

type API interface {
	HandleServerMux(mux *http.ServeMux)
}
