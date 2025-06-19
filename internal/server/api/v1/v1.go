package v1

import (
	"net/http"

	"github.com/kotyurgini/WorkmateTestWork/internal/storage"
)

type APIv1 struct {
	st storage.Storage
}

func NewAPIv1(st storage.Storage) *APIv1 {
	return &APIv1{
		st: st,
	}
}

func (api *APIv1) HandleServerMux(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/task/new", func(w http.ResponseWriter, r *http.Request) {
		api.newTask(w, r)
	})
	mux.HandleFunc("/api/v1/task/get", func(w http.ResponseWriter, r *http.Request) {
		api.getTask(w, r)
	})
	mux.HandleFunc("/api/v1/task/delete", func(w http.ResponseWriter, r *http.Request) {
		api.deleteTask(w, r)
	})
}
