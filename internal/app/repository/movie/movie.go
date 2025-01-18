package movie

import (
	"context"
	"strings"
	"vote-system/database"
	"vote-system/internal/entity"
	"vote-system/pkg/helper"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, movie *entity.Movie) (err error)
	CheckMovieTitleExists(ctx context.Context, request *entity.Movie) (status bool, err error)
	UpdateMovie(ctx context.Context, request *entity.Movie) (err error)
	GetMovieByID(ctx context.Context, request *entity.Movie) (resp *entity.Movie, err error)
}

type movieRepository struct {
	Database *database.Database
}

func NewMovieRepository(database *database.Database) MovieRepository {
	return &movieRepository{
		Database: database,
	}
}

func (mr *movieRepository) CreateMovie(ctx context.Context, movie *entity.Movie) (err error) {
	err = mr.Database.DB.WithContext(ctx).Create(&movie).Error
	if err != nil {
		return err
	}
	return

}

func (mr *movieRepository) CheckMovieTitleExists(ctx context.Context, request *entity.Movie) (status bool, err error) {
	var count int64
	err = mr.Database.DB.WithContext(ctx).
		Model(&entity.Movie{}).
		Where("lower(title) = ? ", strings.ToLower(request.Title)).
		Count(&count).Error

	if err != nil {

		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (mr *movieRepository) GetMovieByID(ctx context.Context, request *entity.Movie) (resp *entity.Movie, err error) {

	err = mr.Database.DB.WithContext(ctx).Where("id = ?", request.ID).Find(&resp).Error
	if err != nil {
		err = helper.HandleError(err)
		return resp, err
	}

	return
}

func (mr *movieRepository) UpdateMovie(ctx context.Context, request *entity.Movie) (err error) {

	err = mr.Database.DB.WithContext(ctx).Save(&request).Error
	if err != nil {
		return err
	}

	return
}
