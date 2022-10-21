package model

type Body struct {
	HasNext  bool  `json:"has_next"`
	PageID   int32 `from:"page_id" binding:"required,min=1"`
	PageSize int32 `from:"page_size" binding:"required,min=5,max=10"`
	Total    int   `json:"total"`
	Result   any   `json:"result"`
}
