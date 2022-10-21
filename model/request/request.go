package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	PageID   int32 `json:"page_id" form:"page_id" binding:"required,min=1"`       // 页码
	PageSize int32 `json:"page_size" form:"page_size" binding:"required,max=100"` // 每页大小
}

// GetById Find by id structure
type GetById struct {
	ID int64 `json:"id" form:"id" binding:"required"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
