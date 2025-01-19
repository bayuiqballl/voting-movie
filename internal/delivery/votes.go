package delivery

import (
	"net/http"
	"strconv"
	"vote-system/internal/app/usecase/votes"
	"vote-system/internal/entity"
	"vote-system/pkg/constant"
	"vote-system/pkg/helper"

	"github.com/gofiber/fiber/v2"
)

type VotesHandler interface {
	Vote(c *fiber.Ctx) (err error)
}

type votesHandler struct {
	service votes.Service
}

func NewVotesHandler(service votes.Service) VotesHandler {
	return &votesHandler{
		service: service,
	}
}

func (vh *votesHandler) Vote(c *fiber.Ctx) (err error) {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	valueCtx := helper.GetValueContext(ctx)

	request := new(entity.VoteRequest)
	if err = c.BodyParser(request); err != nil {
		err = helper.Error(http.StatusBadRequest, constant.MsgInvalidRequest, err)
		return
	}

	idStr, err := strconv.Atoi(valueCtx.UserId)
	if err != nil {
		err = helper.Error(http.StatusBadRequest, constant.MsgInvalidRequest, err)
		return
	}

	request.UserID = idStr
	err = vh.service.UpsertVotes(ctx, request)
	if err != nil {
		return helper.ResponseError(c, err)
	}

	return helper.ResponseCreatedOK(c, constant.Success, nil)
}
