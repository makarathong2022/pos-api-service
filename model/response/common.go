package response

type PageResult struct {
	Result   interface{} `json:"result"`
	HasNext  bool        `json:"has_next"`
	Total    int         `json:"total"`
	Page     int32       `json:"page"`
	PageSize int32       `json:"pageSize"`
}
