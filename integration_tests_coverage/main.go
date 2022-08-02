package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	runServer()
}

func runServer() {
	ctx, cancel := context.WithCancel(context.Background())

	killServer := newKillServer(":8080", cancel)
	go killServer.Start()

	<-ctx.Done()

	killServer.server.Shutdown(context.Background())
}

type killServer struct {
	server http.Server
	cancel context.CancelFunc
}

func newKillServer(addr string, cancel context.CancelFunc) *killServer {
	return &killServer{
		server: http.Server{
			Addr: addr,
		},
		cancel: cancel,
	}
}

func (s *killServer) Start() {
	s.server.Handler = s

	err := s.server.ListenAndServe()
	if err != nil {
		fmt.Println("KillServer Error:", err)
	}

	fmt.Println("KillServer Started")
}

func (s *killServer) Cancel() {
	s.server.Shutdown(context.Background())
	fmt.Println("KillServer Cancelled")
}

func (s *killServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Println("handler called")
	// cancel the context
	s.cancel()
}
