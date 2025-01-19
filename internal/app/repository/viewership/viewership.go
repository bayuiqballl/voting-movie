package viewership

import (
	"context"
	"time"
	"vote-system/database"
	"vote-system/internal/entity"

	"gorm.io/gorm/clause"
)

type ViewershipRepository interface {
	UpsertViewership(ctx context.Context, viewership *entity.Viewership) (err error)
}

type viewershipRepository struct {
	Database *database.Database
}

func NewViewershipRepository(database *database.Database) ViewershipRepository {
	return &viewershipRepository{
		Database: database,
	}
}

func (vr *viewershipRepository) UpsertViewership(ctx context.Context, viewership *entity.Viewership) (err error) {
	// Perform an upsert
	if watchedAt := viewership.WatchedAt; watchedAt.IsZero() {
		viewership.WatchedAt = time.Now()
	}

	err = vr.Database.DB.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "movie_id"}},       // Define the conflict columns
			DoUpdates: clause.AssignmentColumns([]string{"duration", "watched_at"}), // Update these fields
		}).
		Create(&viewership).Error
	if err != nil {
		return err
	}

	return nil
}
