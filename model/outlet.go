package model

type Outlet struct {
	OutletCd    int32  `json:"outlet_cd" binding:"required"`
	OutletName  string `json:"outlet_name" binding:"required"`
	IpAddress   string `json:"ip_address" binding:"required"`
	Sort        int64  `json:"sort" binding:"required"`
	Description string `json:"description"`
}
