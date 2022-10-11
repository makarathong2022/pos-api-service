package model

type MenuItemDetail struct {
	ItemCdDetail int32  `json:"item_cd_detail"  binding:"required"`
	ItemID       int32  `json:"item_id" binding:"required"`
	GroupID      int32  `json:"group_id" binding:"required"`
	CategoryID   int32  `json:"category_id" binding:"required"`
	SizeID       int32  `json:"size_id" binding:"required"`
	Cost         string `json:"cost" binding:"required"`
	Price        string `json:"price" binding:"required"`
	VatID        int32  `json:"vat_id" binding:"required"`
	Vat          string `json:"vat"  binding:"required"`
	TerminalID   int32  `json:"terminal_id" binding:"required"`
	OutletID     int32  `json:"outlet_id" binding:"required"`
	PrinterID    int32  `json:"printer_id" binding:"required"`
}
