package delivery

import (
	"net/http"
	"vote-system/internal/app/usecase/admin"
	"vote-system/internal/app/usecase/movie"
	"vote-system/internal/entity"
	"vote-system/pkg/constant"
	"vote-system/pkg/helper"

	"github.com/gofiber/fiber/v2"
)

type AdminHandler interface {
	RegisterAdmin(c *fiber.Ctx) (err error)
	LoginAdmin(c *fiber.Ctx) (err error)
	UploadMovie(c *fiber.Ctx) (err error)
}

type adminHandler struct {
	service      admin.Service
	movieService movie.Service
}

func NewAdminHandler(service admin.Service, movieService movie.Service) AdminHandler {
	return &adminHandler{
		service:      service,
		movieService: movieService,
	}
}

func (h *adminHandler) RegisterAdmin(c *fiber.Ctx) (err error) {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	request := new(entity.Admin)
	if err = c.BodyParser(request); err != nil {
		err = helper.Error(http.StatusBadRequest, constant.MsgInvalidRequest, err)
		return
	}

	err = h.service.RegisterAdmin(ctx, request)
	if err != nil {
		return helper.ResponseError(c, err)
	}

	return helper.ResponseCreatedOK(c, constant.Success, nil)
}

func (h *adminHandler) LoginAdmin(c *fiber.Ctx) (err error) {

	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	request := new(entity.LoginAdmin)
	if err = c.BodyParser(request); err != nil {
		err = helper.Error(http.StatusBadRequest, constant.MsgInvalidRequest, err)
		return
	}

	resp, err := h.service.LoginAdmin(ctx, request)
	if err != nil {
		return helper.ResponseError(c, err)
	}

	return helper.ResponseOK(c, constant.Success, resp)

}

func (h *adminHandler) UploadMovie(c *fiber.Ctx) (err error) {
	_, cancel := helper.CreateContextWithTimeout()
	defer cancel()

	file, err := c.FormFile("file")
	if err != nil {
		err = helper.Error(http.StatusBadRequest, constant.MsgInvalidRequest, err)
		return
	}

	resp, err := h.movieService.UploadMovie(c, file)
	if err != nil {
		return helper.ResponseError(c, err)
	}

	return helper.ResponseCreatedOK(c, constant.Success, resp)

}
