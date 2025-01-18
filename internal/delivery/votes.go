package delivery

import (
	"vote-system/internal/app/usecase/votes"
)

type VotesHandler interface{}

type votesHandler struct {
	service votes.Service
}

func NewVotesHandler(service votes.Service) VotesHandler {
	return &votesHandler{
		service: service,
	}
}
