package admin

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"vote-system/cmd/middleware"
	"vote-system/internal/entity"
	"vote-system/pkg/constant"
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

	encryptedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return err
	}

	request.Password = encryptedPassword

	err = au.repository.CreateAdmin(ctx, request)
	if err != nil {
		return
	}

	return
}

func (au *service) LoginAdmin(ctx context.Context, request *entity.LoginAdmin) (resp *entity.LoginAdminResponse, err error) {

	if err := au.validator.Validate(request); err != nil {
		err = helper.Error(http.StatusBadRequest, err.Error(), err)
		return nil, err
	}

	getAdmin, err := au.repository.GetAdminByEmail(ctx, &entity.Admin{
		Email: request.Email,
	})

	if err != nil {
		return nil, err
	}

	if !helper.CheckPassword(getAdmin.Password, request.Password) {
		err = helper.Error(http.StatusUnauthorized, constant.MsgWrongPassword, nil)
		return nil, err
	}

	expirationTime := time.Now().Add(time.Hour * time.Duration(24))

	IdStr := strconv.Itoa(getAdmin.ID)
	token, err := middleware.GenerateToken(IdStr, getAdmin.Email, getAdmin.Role)
	if err != nil {
		return nil, err
	}

	resp = &entity.LoginAdminResponse{
		Token:     token,
		ExpiredAt: expirationTime,
	}

	return
}
