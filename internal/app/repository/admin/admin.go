package admin

import (
	"context"
	"strings"
	"vote-system/database"
	"vote-system/internal/entity"
)

type AdminRepository interface {
	CreateAdmin(ctx context.Context, admin *entity.Admin) (err error)
	CheckAccountNameExists(ctx context.Context, admin *entity.Admin) (status bool, err error)
	GetAdmin(ctx context.Context, admin *entity.Admin) (err error)
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

func (ar *adminRepository) GetAdmin(ctx context.Context, admin *entity.Admin) (err error) {
	err = ar.Database.DB.WithContext(ctx).First(&admin).Error
	if err != nil {
		return err
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
