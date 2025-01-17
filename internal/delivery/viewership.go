package delivery

import "vote-system/internal/app/usecase/movie"

type ViewershipHandler struct {
	service movie.Service
}

func NewViewershipHandler(service movie.Service) ViewershipHandler {
	return ViewershipHandler{
		service: service,
	}
}
