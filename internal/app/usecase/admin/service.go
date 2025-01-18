package admin

import (
	"context"
	repository "vote-system/internal/app/repository/admin"
	"vote-system/internal/entity"
	"vote-system/pkg/identifier"
	"vote-system/pkg/validator"
)

type Service interface {
	RegisterAdmin(ctx context.Context, request *entity.Admin) (err error)
	LoginAdmin(ctx context.Context, request *entity.LoginAdmin) (resp *entity.LoginAdminResponse, err error)
}

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
