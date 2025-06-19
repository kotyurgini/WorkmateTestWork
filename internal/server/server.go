package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kotyurgini/WorkmateTestWork/internal/server/api"
	v1 "github.com/kotyurgini/WorkmateTestWork/internal/server/api/v1"
	"github.com/kotyurgini/WorkmateTestWork/internal/storage"
)

type Server struct {
	serv *http.Server
	api  api.API
	st   storage.Storage
}

func NewServer(st storage.Storage, port int) *Server {
	mux := http.NewServeMux()

	api := v1.NewAPIv1(st)
	api.HandleServerMux(mux)

	return &Server{
		serv: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		api: api,
		st:  st,
	}
}

func (s *Server) Start() {
	go func() {
		log.Printf("Server start on http://localhost%s\n", s.serv.Addr)
		err := s.serv.ListenAndServe()
		if err != nil {
			log.Println(err.Error())
		}
	}()
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := s.serv.Shutdown(ctx); err != nil {
		log.Println(err.Error())
	}
	s.st.Close()
}
