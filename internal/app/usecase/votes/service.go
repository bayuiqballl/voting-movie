package votes

import (
	"context"
	repository "vote-system/internal/app/repository/votes"
	"vote-system/internal/entity"
	"vote-system/pkg/identifier"
	"vote-system/pkg/validator"
)

type Service interface {
	UpsertVotes(ctx context.Context, request *entity.VoteRequest) (err error)
	GetListUserVotes(ctx context.Context, movieID int) (votes []entity.GetListUserVoteResponse, err error)
}

type service struct {
	repository repository.VotesRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewService(repository repository.VotesRepository, validator validator.Validator, identifier identifier.Identifier) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}
