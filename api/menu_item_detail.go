package api

import (
	"net/http"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/gin-gonic/gin"
)

func (server *Server) createMenuItemDetail(ctx *gin.Context) {
	var body model.MenuItemDetail

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMenuItemDetailParams{
		ItemCdDetail: body.ItemCdDetail,
		ItemID:       body.ItemID,
		GroupID:      body.GroupID,
		CategoryID:   body.CategoryID,
		SizeID:       body.SizeID,
		Cost:         body.Cost,
		Price:        body.Price,
		OutletID:     body.OutletID,
		PrinterID:    body.PrinterID,
		VatID:        body.VatID,
		Vat:          body.Vat,
		TerminalID:   body.TerminalID,
	}

	menuItemDetail, err := server.store.CreateMenuItemDetail(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, menuItemDetail)
}
