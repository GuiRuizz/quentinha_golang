package response

type PaginatedProductResponse struct {
	Data       []ProductResponse `json:"data"`
	Page       int64             `json:"page"`
	Limit      int64             `json:"limit"`
	Total      int64             `json:"total"`
	TotalPages int64             `json:"totalPages"`
}
