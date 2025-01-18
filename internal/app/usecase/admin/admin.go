package admin

import (
	"context"
	"net/http"
	"vote-system/internal/entity"
	"vote-system/pkg/helper"
)

func (au *service) RegisterAdmin(ctx context.Context, request *entity.Admin) (err error) {
	if err = au.validator.Validate(request); err != nil {
		err = helper.Error(http.StatusBadRequest, "validation error", err)
		return
	}

	if request.Role != "admin" {
		err = helper.Error(http.StatusBadRequest, "role must be admin", nil)
		return
	}

	status, err := au.repository.CheckAccountNameExists(ctx, request)
	if err != nil {
		return err
	}

	if status {
		err = helper.Error(http.StatusBadRequest, "email already exists", nil)
		return
	}

	request.Password = helper.EncryptPassword(request.Password)

	err = au.repository.CreateAdmin(ctx, request)
	if err != nil {
		return
	}

	return
}
