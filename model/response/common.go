package response

type PageResult struct {
	List     interface{} `json:"list"`
	HasNext  bool        `json:"has_next"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
