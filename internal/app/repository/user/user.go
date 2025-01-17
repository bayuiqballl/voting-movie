package user

import "vote-system/database"

type UserRepository interface {
}

type userRepository struct {
	Database *database.Database
}

func NewUserRepository(database *database.Database) UserRepository {
	return &userRepository{
		Database: database,
	}
}
