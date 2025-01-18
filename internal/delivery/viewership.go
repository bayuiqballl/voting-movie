package delivery

import (
	"vote-system/internal/app/usecase/viewership"
)

type ViewershipHandler struct {
	service viewership.Service
}

func NewViewershipHandler(service viewership.Service) ViewershipHandler {
	return ViewershipHandler{
		service: service,
	}
}
