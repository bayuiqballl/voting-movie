package user

import (
	"context"
	repository "vote-system/internal/app/repository/user"
	"vote-system/internal/entity"
	"vote-system/pkg/identifier"
	"vote-system/pkg/validator"
)

type Service interface {
	CreateUser(ctx context.Context, request *entity.User) (err error)
	LoginUser(ctx context.Context, request *entity.LoginUser) (resp *entity.LoginUserResponse, err error)
}

type service struct {
	repository repository.UserRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewService(repository repository.UserRepository, validator validator.Validator, identifier identifier.Identifier) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}
