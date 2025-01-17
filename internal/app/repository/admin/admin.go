package admin

import "vote-system/database"

type AdminRepository interface {
}

type adminRepository struct {
	Database *database.Database
}

func NewAdminRepository(database *database.Database) AdminRepository {
	return &adminRepository{
		Database: database,
	}
}
