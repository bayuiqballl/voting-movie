package votes

import (
	"context"
	"vote-system/database"
	"vote-system/internal/entity"
)

type VotesRepository interface {
	CreateVotes(ctx context.Context, votes *entity.Vote) (err error)
	DeleteVotes(ctx context.Context, votes *entity.Vote) (err error)
	CheckVoteExists(ctx context.Context, votes *entity.Vote) (status bool, err error)
}

type votesRepository struct {
	Database *database.Database
}

func NewVotesRepository(database *database.Database) VotesRepository {
	return &votesRepository{
		Database: database,
	}
}

func (vr *votesRepository) CreateVotes(ctx context.Context, votes *entity.Vote) (err error) {
	err = vr.Database.DB.WithContext(ctx).Create(&votes).Error
	if err != nil {
		return err
	}

	return
}

func (vr *votesRepository) DeleteVotes(ctx context.Context, votes *entity.Vote) (err error) {
	// Use Where before calling Delete to construct the query
	err = vr.Database.DB.WithContext(ctx).
		Where("movie_id = ? AND user_id = ?", votes.MovieID, votes.UserID).
		Delete(&entity.Vote{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (vr *votesRepository) CheckVoteExists(ctx context.Context, votes *entity.Vote) (status bool, err error) {
	var count int64
	err = vr.Database.DB.WithContext(ctx).
		Model(&entity.Vote{}).
		Where("user_id = ? AND movie_id = ?", votes.UserID, votes.MovieID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
