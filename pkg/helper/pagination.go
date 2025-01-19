package helper

type PaginationRequest struct {
	Page  int `query:"page" validate:"required"`
	Limit int `query:"limit" validate:"required"`
}

type PaginatedResponse struct {
	TotalCount int64       `json:"total_count"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Data       interface{} `json:"data"`
}

func NewPaginatedResponse(page, limit int, totalCount int64, data interface{}) PaginatedResponse {
	return PaginatedResponse{
		TotalCount: totalCount,
		Page:       page,
		Limit:      limit,
		Data:       data,
	}
}
