package api

import (
	"net/http"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/gin-gonic/gin"
)

func (server *Server) createMenuItemDetail(ctx *gin.Context) {
	var req model.MenuItemDetail

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMenuItemDetailParams{
		ItemCdDetail: req.ItemCdDetail,
		ItemID:       req.ItemID,
		GroupID:      req.GroupID,
		CategoryID:   req.CategoryID,
		SizeID:       req.SizeID,
		Cost:         req.Cost,
		Price:        req.Price,
		OutletID:     req.OutletID,
		PrinterID:    req.PrinterID,
		VatID:        req.VatID,
		Vat:          req.Vat,
		TerminalID:   req.TerminalID,
	}

	menuItemDetail, err := server.store.CreateMenuItemDetail(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, menuItemDetail)
}
