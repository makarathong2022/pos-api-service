package model

import "time"

type MenuModify struct {
	ModifyCd    int32     `json:"modify_cd" binding:"required"`
	ModifyName  string    `json:"modify_name" binding:"required"`
	Sort        int64     `json:"sort" binding:"required"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
