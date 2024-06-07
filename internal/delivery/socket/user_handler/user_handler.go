package user_handler

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/user"
	socketio "github.com/googollee/go-socket.io"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type UserSocketHandler struct {
	Log         *logrus.Logger
	UserUsecase user.UseCase
}

func NewUserSocketHandler(log *logrus.Logger, userUsecase user.UseCase) *UserSocketHandler {
	return &UserSocketHandler{
		Log:         log,
		UserUsecase: userUsecase,
	}
}

func (h *UserSocketHandler) CountAllUser(s socketio.Conn) error {
	ctx := context.Background()
	for {
		count, err := h.UserUsecase.CountAllUser(ctx)
		if err != nil {
			s.Emit("count", "error: "+err.Error())
		} else {
			s.Emit("count", "user count: "+strconv.Itoa(int(count)))
		}

		time.Sleep(2 * time.Second)
	}
}
