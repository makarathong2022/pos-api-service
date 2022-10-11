package model

type Category struct {
	CategoryCd   int32  `json:"category_cd" binding:"required"`
	CategoryName string `json:"category_name" binding:"required"`
	Sort         int64  `json:"sort" binding:"required"`
	Description  string `json:"description"`
}
