// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type TblMenuCategory struct {
	ID           int64     `json:"id"`
	CategoryCd   int32     `json:"category_cd"`
	CategoryName string    `json:"category_name"`
	Sort         int64     `json:"sort"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type TblMenuGroup struct {
	ID          int64     `json:"id"`
	GroupCd     int32     `json:"group_cd"`
	GroupName   string    `json:"group_name"`
	Sort        int64     `json:"sort"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type TblMenuItem struct {
	ID          int64     `json:"id"`
	ItemCd      int32     `json:"item_cd"`
	ItemName    string    `json:"item_name"`
	Abv         string    `json:"abv"`
	Sort        int64     `json:"sort"`
	OutletID    int32     `json:"outlet_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type TblMenuItemDetail struct {
	ID           int64 `json:"id"`
	ItemCdDetail int32 `json:"item_cd_detail"`
	ItemID       int32 `json:"item_id"`
	GroupID      int32 `json:"group_id"`
	CategoryID   int32 `json:"category_id"`
	SizeID       int32 `json:"size_id"`
	// must positive
	Cost string `json:"cost"`
	// must be positive
	Price      string    `json:"price"`
	VatID      int32     `json:"vat_id"`
	Vat        string    `json:"vat"`
	TerminalID int32     `json:"terminal_id"`
	OutletID   int32     `json:"outlet_id"`
	PrinterID  int32     `json:"printer_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type TblMenuItemModify struct {
	ID       int64 `json:"id"`
	ItemCd   int32 `json:"item_cd"`
	ModifyCd int32 `json:"modify_cd"`
}

type TblMenuModify struct {
	ID          int64        `json:"id"`
	ModifyCd    int32        `json:"modify_cd"`
	ModifyName  string       `json:"modify_name"`
	Sort        int64        `json:"sort"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}

type TblMenuSize struct {
	ID          int64     `json:"id"`
	SizeCd      int32     `json:"size_cd"`
	SizeName    string    `json:"size_name"`
	Sort        int64     `json:"sort"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type TblOutlet struct {
	ID          int64     `json:"id"`
	OutletCd    int32     `json:"outlet_cd"`
	OutletName  string    `json:"outlet_name"`
	IpAddress   string    `json:"ip_address"`
	Sort        int64     `json:"sort"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type TblPrinter struct {
	ID          int64     `json:"id"`
	PrintCd     int32     `json:"print_cd"`
	PrintName   string    `json:"print_name"`
	Sort        int64     `json:"sort"`
	IpAddress   string    `json:"ip_address"`
	Description string    `json:"description"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type TblTerminal struct {
	ID           int64     `json:"id"`
	TerminalCd   int32     `json:"terminal_cd"`
	TerminalName string    `json:"terminal_name"`
	IpAddress    string    `json:"ip_address"`
	Sort         int64     `json:"sort"`
	OutletCd     int32     `json:"outlet_cd"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type TblVat struct {
	ID          int64     `json:"id"`
	VatCd       int32     `json:"vat_cd"`
	VatKey      string    `json:"vat_key"`
	VatName     string    `json:"vat_name"`
	Sort        int64     `json:"sort"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
