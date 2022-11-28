package model

type Tax struct {
	VatCd       int32  `json:"vat_cd" binding:"required"`
	VatKey      string `json:"vat_key" binding:"required"`
	VatName      string `json:"vat_name" binding:"required"`
	Sort        int64  `json:"sort" binding:"required"`
	Description string `json:"description"`
}
