package movie

import (
	"context"
	"mime/multipart"
	repository "vote-system/internal/app/repository/movie"
	"vote-system/internal/entity"
	"vote-system/pkg/helper"
	"vote-system/pkg/identifier"
	"vote-system/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	UploadMovie(c *fiber.Ctx, file *multipart.FileHeader) (resp entity.UploadMovieResponse, err error)
	InsertMovie(context context.Context, request *entity.Movie) (err error)
	UpdateMovie(context context.Context, request *entity.Movie) (err error)
	GetListMovies(context context.Context, request entity.GetListMovieRequest) (resp helper.PaginatedResponse, err error)
	GetMostDataMovie(context context.Context) (resp *entity.GetMostDataMovie, err error)
}

type service struct {
	repository repository.MovieRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewService(repository repository.MovieRepository, validator validator.Validator, identifier identifier.Identifier) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}
