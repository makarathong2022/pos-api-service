package model

type Printer struct {
	PrinterCd   int32  `json:"print_cd" binding:"required"`
	PrinterName string `json:"print_name" binding:"required"`
	Sort        int64  `json:"sort" binding:"required"`
	IpAddress   string `json:"ip_address" binding:"required"`
	Description string `json:"description"`
}
