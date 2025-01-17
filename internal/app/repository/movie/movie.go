package movie

import "vote-system/database"

type MovieRepository interface {
}

type movieRepository struct {
	Database *database.Database
}

func NewMovieRepository(database *database.Database) MovieRepository {
	return &movieRepository{
		Database: database,
	}
}
