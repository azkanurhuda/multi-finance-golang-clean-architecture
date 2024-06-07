package socket

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/socket/user_handler"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"time"
)

type SocketHandler struct {
	Server            *socketio.Server
	App               *fiber.App
	UserSocketHandler *user_handler.UserSocketHandler
	User              user.UseCase
}

func (h *SocketHandler) RegisterHandlers() {
	h.Server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	h.Server.OnEvent("/pengguna", "message", func(s socketio.Conn) error {
		//dat_a, _ := h.User.CountAllUser(context.Background())
		for {
			//fmt.Println(s.ID())
			s.Emit("message", "hallo")
			time.Sleep(1 * time.Second)
		}
		return nil
	})

	h.Server.OnEvent("/user", "count", h.UserSocketHandler.CountAllUser)

	h.Server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("disconnected:", s.ID(), reason)
	})
}
