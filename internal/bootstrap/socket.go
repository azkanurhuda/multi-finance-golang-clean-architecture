package bootstrap

import (
	"fmt"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/config"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func NewSocketIOServer(cfg *config.Config) *socketio.Server {
	server := socketio.NewServer(nil)

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("SocketIO listen error: %s\n", err)
		}
	}()

	http.Handle("/socket.io/", server)

	go func() {
		address := fmt.Sprintf(":%d", cfg.Server.SocketPort)
		log.Printf("Serving Socket.IO at localhost%s\n", address)
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Fatalf("HTTP server listen error: %s\n", err)
		}
	}()

	return server
}
