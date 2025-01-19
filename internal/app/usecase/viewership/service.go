package viewership

import (
	"context"
	repository "vote-system/internal/app/repository/viewership"
	"vote-system/internal/entity"
	"vote-system/pkg/identifier"
	"vote-system/pkg/validator"
)

type Service interface {
	UpsertViewership(ctx context.Context, viewership *entity.Viewership) (err error)
}

type service struct {
	repository repository.ViewershipRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewService(repository repository.ViewershipRepository, validator validator.Validator, identifier identifier.Identifier) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}
