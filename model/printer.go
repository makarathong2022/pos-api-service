package model

type Printer struct {
	PrinterCd   int32  `json:"print_cd" binding:"required"`
	PrinterName string `json:"print_name" binding:"required"`
	Sort        int64  `json:"sort" binding:"required"`
	IpAddress   string `json:"ip_address" binding:"required"`
	Description string `json:"description"`
}

// AKIAWDUKBNEY6FSWBD6B
// zjKReHQF3X2DSpdcS2oQXWa1WX6ZpRms0yfRI/LT
// Database login
// username: root
// Master Password: moHoAwOyqbdN11ha46fe
// postgresql://pos-api-service.ci5bs4qfk9kq.ap-southeast-1.rds.amazonaws.com/outlet_0001
