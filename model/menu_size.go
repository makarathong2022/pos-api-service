package model

type MenuSize struct {
	SizeCd      int32  `json:"size_cd" binding:"required"`
	SizeName    string `json:"size_name" binding:"required"`
	Sort        int64  `json:"sort" binding:"required"`
	Description string `json:"description"`
}
