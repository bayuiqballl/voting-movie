package delivery

import (
	"net/http"
	"vote-system/internal/app/usecase/viewership"
	"vote-system/internal/entity"
	"vote-system/pkg/constant"
	"vote-system/pkg/helper"

	"github.com/gofiber/fiber/v2"
)

type ViewershipHandler interface {
	UpsertViewership(c *fiber.Ctx) (err error)
}

type viewershipHandler struct {
	service viewership.Service
}

func NewViewershipHandler(service viewership.Service) ViewershipHandler {
	return &viewershipHandler{
		service: service,
	}
}

func (h *viewershipHandler) UpsertViewership(c *fiber.Ctx) (err error) {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	request := new(entity.Viewership)
	if err = c.BodyParser(request); err != nil {
		err = helper.Error(http.StatusBadRequest, constant.MsgInvalidRequest, err)
		return
	}

	err = h.service.UpsertViewership(ctx, request)
	if err != nil {
		return helper.ResponseError(c, err)
	}

	return helper.ResponseCreatedOK(c, constant.Success, nil)
}
