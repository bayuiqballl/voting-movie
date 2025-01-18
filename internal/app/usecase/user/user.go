package user

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

func (s *service) CreateUser(ctx context.Context, request *entity.User) (err error) {

	if err := s.validator.Validate(request); err != nil {
		err = helper.Error(http.StatusBadRequest, err.Error(), err)
		return err
	}

	status, err := s.repository.CheckUserExists(ctx, request)
	if err != nil {
		return err
	}

	if status {
		err = helper.Error(http.StatusBadRequest, "email already exists", nil)
		return err
	}

	encryptedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return err
	}

	request.Password = encryptedPassword

	err = s.repository.CreateUser(ctx, request)
	if err != nil {
		return err
	}
	return
}

func (s *service) LoginUser(ctx context.Context, request *entity.LoginUser) (resp *entity.LoginUserResponse, err error) {

	if err := s.validator.Validate(request); err != nil {
		err = helper.Error(http.StatusBadRequest, err.Error(), err)
		return nil, err
	}

	getUser, err := s.repository.GetUserByEmail(ctx, &entity.User{
		Email: request.Email,
	})

	if err != nil {
		return nil, err
	}

	if !helper.CheckPassword(getUser.Password, request.Password) {
		err = helper.Error(http.StatusUnauthorized, constant.MsgWrongPassword, nil)
		return nil, err
	}

	expirationTime := time.Now().Add(time.Hour * time.Duration(24))

	IdStr := strconv.Itoa(getUser.ID)
	token, err := middleware.GenerateToken(IdStr, getUser.Email, "")
	if err != nil {
		return nil, err
	}

	resp = &entity.LoginUserResponse{
		Token:     token,
		ExpiredAt: expirationTime,
	}

	return

}
