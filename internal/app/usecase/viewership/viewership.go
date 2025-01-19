package viewership

import (
	"context"
	"net/http"
	"vote-system/internal/entity"
	"vote-system/pkg/helper"
)

func (s *service) UpsertViewership(ctx context.Context, request *entity.Viewership) (err error) {
	if err := s.validator.Validate(request); err != nil {
		err = helper.Error(http.StatusBadRequest, err.Error(), err)
		return err
	}

	getMovieDetail, err := s.movieRepo.GetMovieByID(ctx, &entity.Movie{
		ID: request.MovieID,
	})
	if err != nil {
		return err
	}

	if getMovieDetail.ID == 0 {
		err = helper.Error(http.StatusBadRequest, "movie not found", nil)
		return err
	}

	if getMovieDetail.Duration >= request.Duration {
		err = s.repository.UpsertViewership(ctx, request)
		if err != nil {
			return err
		}
	}

	return
}
