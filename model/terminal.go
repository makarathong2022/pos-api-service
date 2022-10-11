package model

type Terminal struct {
	TerminalCd   int32  `json:"terminal_cd" binding:"required"`
	TerminalName string `json:"terminal_name" binding:"required"`
	IpAddress    string `json:"ip_address" binding:"required"`
	Sort         int64  `json:"sort" binding:"required"`
	OutletCd     int32  `json:"outlet_cd" binding:"required"`
	Description  string `json:"description"`
}
