package delivery

import "vote-system/internal/app/usecase/admin"

type AdminHandler interface{}

type adminHandler struct {
	service admin.Service
}

func NewAdminHandler(service admin.Service) AdminHandler {
	return &adminHandler{
		service: service,
	}
}
