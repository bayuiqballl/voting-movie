package votes

import (
	"context"
	"net/http"
	"vote-system/internal/entity"
	"vote-system/pkg/helper"
)

func (vs *service) UpsertVotes(ctx context.Context, request *entity.VoteRequest) (err error) {

	if err := vs.validator.Validate(request); err != nil {
		err = helper.Error(http.StatusBadRequest, err.Error(), err)
		return err
	}

	var vote entity.Vote
	vote.UserID = request.UserID
	vote.MovieID = request.MovieID

	status, err := vs.repository.CheckVoteExists(ctx, &vote)
	if err != nil {
		return err
	}
	if request.IsVote {
		if status {
			err = helper.Error(http.StatusBadRequest, "vote already exists", nil)
			return err
		}

		err = vs.repository.CreateVotes(ctx, &vote)
		if err != nil {
			return err
		}
	}

	if !request.IsVote {
		err = vs.repository.DeleteVotes(ctx, &vote)
		if err != nil {
			return err
		}
	}

	return

}

func (vs *service) GetListUserVotes(ctx context.Context, movieID int) (votes []entity.GetListUserVoteResponse, err error) {

	if movieID == 0 {
		err = helper.Error(http.StatusBadRequest, "movie ID is required", nil)
		return nil, err
	}

	votes, err = vs.repository.GetListUserVotes(ctx, movieID)
	if err != nil {
		return nil, err
	}

	return

}
