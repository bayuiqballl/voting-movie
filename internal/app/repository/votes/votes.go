package votes

import "vote-system/database"

type VotesRepository interface{}

type votesRepository struct {
	Database *database.Database
}

func NewVotesRepository(database *database.Database) VotesRepository {
	return &votesRepository{
		Database: database,
	}
}
