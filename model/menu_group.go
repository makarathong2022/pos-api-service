package model

type MenuGroup struct {
	GroupCd     int32  `json:"group_cd" binding:"required"`
	GroupName   string `json:"group_name" binding:"required"`
	Sort        int64  `json:"sort" binding:"required"`
	Description string `json:"description"`
}
