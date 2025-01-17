package delivery

import "vote-system/internal/app/usecase/movie"

type UserHandler struct {
	service movie.Service
}

func NewUserHandler(service movie.Service) UserHandler {
	return UserHandler{
		service: service,
	}
}
