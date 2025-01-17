package delivery

import "vote-system/internal/app/usecase/movie"

type MovieHandler interface{}

type movieHandler struct {
	service movie.Service
}

func NewMovieHandler(service movie.Service) MovieHandler {
	return &movieHandler{
		service: service,
	}
}
