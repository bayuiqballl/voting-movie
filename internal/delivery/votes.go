package delivery

import (
	"vote-system/internal/app/usecase/movie"
	"vote-system/internal/app/usecase/votes"
)

type VotesHandler interface{}

type votesHandler struct {
	service votes.Service
}

func NewVotesHandler(service movie.Service) VotesHandler {
	return &votesHandler{
		service: service,
	}
}
