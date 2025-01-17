package viewership

import "vote-system/database"

type ViewershipRepository interface{}

type viewershipRepository struct {
	Database *database.Database
}

func NewViewershipRepository(database *database.Database) ViewershipRepository {
	return &viewershipRepository{
		Database: database,
	}
}
