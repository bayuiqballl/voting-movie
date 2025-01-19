package delivery

import (
	"net/http"
	"vote-system/internal/app/usecase/movie"
	"vote-system/internal/entity"
	"vote-system/pkg/constant"
	"vote-system/pkg/helper"

	"github.com/gofiber/fiber/v2"
)

type MovieHandler interface {
	GetListMovies(c *fiber.Ctx) (err error)
	GetMostDataMovie(c *fiber.Ctx) (err error)
}

type movieHandler struct {
	service movie.Service
}

func NewMovieHandler(service movie.Service) MovieHandler {
	return &movieHandler{
		service: service,
	}
}

func (h *movieHandler) GetListMovies(c *fiber.Ctx) (err error) {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	request := new(entity.GetListMovieRequest)
	if err = c.QueryParser(request); err != nil {
		err = helper.Error(http.StatusBadRequest, constant.MsgInvalidRequest, err)
		return
	}

	resp, err := h.service.GetListMovies(ctx, *request)
	if err != nil {
		return helper.ResponseError(c, err)
	}

	return helper.ResponseOK(c, constant.Success, resp)

}

func (h *movieHandler) GetMostDataMovie(c *fiber.Ctx) (err error) {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	resp, err := h.service.GetMostDataMovie(ctx)
	if err != nil {
		return helper.ResponseError(c, err)
	}

	return helper.ResponseOK(c, constant.Success, resp)

}
