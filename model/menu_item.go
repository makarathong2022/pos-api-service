package model

import "time"

type MenuItem struct {
	ItemCd      int32     `json:"item_cd" binding:"required"`
	ItemName    string    `json:"item_name" binding:"required"`
	Abv         string    `json:"abv" binding:"required"`
	Sort        int64     `json:"sort" binding:"required"`
	Description string    `json:"description"`
	OutletID    int32     `json:"outlet_id" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
