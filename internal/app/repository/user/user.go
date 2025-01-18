package user

import (
	"context"
	"strings"
	"vote-system/database"
	"vote-system/internal/entity"
	"vote-system/pkg/helper"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (err error)
	CheckUserExists(ctx context.Context, user *entity.User) (status bool, err error)
	GetUserByEmail(ctx context.Context, request *entity.User) (resp *entity.User, err error)
}

type userRepository struct {
	Database *database.Database
}

func NewUserRepository(database *database.Database) UserRepository {
	return &userRepository{
		Database: database,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *entity.User) (err error) {
	err = ur.Database.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return
}

func (ur *userRepository) CheckUserExists(ctx context.Context, user *entity.User) (status bool, err error) {
	var count int64
	err = ur.Database.DB.WithContext(ctx).
		Model(&entity.User{}).
		Where("lower(email) = ? ", strings.ToLower(user.Email)).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil

}

func (ur *userRepository) GetUserByEmail(ctx context.Context, request *entity.User) (resp *entity.User, err error) {
	err = ur.Database.DB.WithContext(ctx).Where("email = ?", request.Email).Find(&resp).Error
	if err != nil {
		err = helper.HandleError(err)
		return resp, err
	}

	return
}
