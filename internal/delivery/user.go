package delivery

import (
	"vote-system/internal/app/usecase/user"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) UserHandler {
	return UserHandler{
		service: service,
	}
}
