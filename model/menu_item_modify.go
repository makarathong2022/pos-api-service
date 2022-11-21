package model

type MenuItemModify struct {
	ItemCd   int32 `json:"item_cd" binding:"required"`
	ModifyCd int32 `json:"modify_cd" binding:"required"`
}
