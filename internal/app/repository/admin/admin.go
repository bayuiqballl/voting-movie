package admin

import (
	"context"
	"strings"
	"vote-system/database"
	"vote-system/internal/entity"
	"vote-system/pkg/helper"
)

type AdminRepository interface {
	CreateAdmin(ctx context.Context, admin *entity.Admin) (err error)
	CheckAccountNameExists(ctx context.Context, admin *entity.Admin) (status bool, err error)
	GetAdminByEmail(ctx context.Context, request *entity.Admin) (resp *entity.Admin, err error)
}

type adminRepository struct {
	Database *database.Database
}

func NewAdminRepository(database *database.Database) AdminRepository {
	return &adminRepository{
		Database: database,
	}
}

func (ar *adminRepository) CreateAdmin(ctx context.Context, admin *entity.Admin) (err error) {
	err = ar.Database.DB.WithContext(ctx).Create(&admin).Error
	if err != nil {
		return err
	}
	return
}

func (ar *adminRepository) GetAdminByEmail(ctx context.Context, request *entity.Admin) (resp *entity.Admin, err error) {
	err = ar.Database.DB.WithContext(ctx).Where("email = ?", request.Email).Find(&resp).Error
	if err != nil {
		err = helper.HandleError(err)
		return resp, err
	}

	return
}

func (ar *adminRepository) CheckAccountNameExists(ctx context.Context, admin *entity.Admin) (status bool, err error) {
	var count int64
	err = ar.Database.DB.WithContext(ctx).
		Model(&entity.Admin{}).
		Where("email = ? ", strings.ToLower(admin.Email)).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
