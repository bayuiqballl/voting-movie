package admin

import (
	repository "vote-system/internal/app/repository/admin"
	"vote-system/pkg/identifier"
	"vote-system/pkg/validator"
)

type Service interface{}

type service struct {
	repository repository.AdminRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewService(repository repository.AdminRepository, validator validator.Validator, identifier identifier.Identifier) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}
