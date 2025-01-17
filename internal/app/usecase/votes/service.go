package votes

import (
	repository "vote-system/internal/app/repository/votes"
	"vote-system/pkg/identifier"
	"vote-system/pkg/validator"
)

type Service interface{}

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
