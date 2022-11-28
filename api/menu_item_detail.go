package api

import (
	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/model/request"
	"github.com/boincompany/pos_api_service/model/response"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) GetMenuItemDetials(ctx *gin.Context) {
	var req request.PageInfo

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.GetMenuItemDetialsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	menuItemDetails, err := server.store.GetMenuItemDetials(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	total := len(menuItemDetails)

	response.OkWithDetailed(response.PageResult{
		Result:   menuItemDetails,
		Page:     req.PageID,
		PageSize: req.PageSize,
		Total:    total,
		HasNext:  total == int(req.PageSize),
	}, utils.GET_SUCCESS, ctx)

}

func (server *Server) GetMenuItemDetial(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	menuItemDetail, err := server.store.GetMenuItemDetial(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithData(menuItemDetail, ctx)

}

func (server *Server) CreateMenuItemDetail(ctx *gin.Context) {
	var body model.MenuItemDetail

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
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
		response.FailWithMessage(errRes(err), ctx)
		return
	}
	response.OkWithData(menuItemDetail, ctx)
}

func (server *Server) UpdateMenuItemDetial(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	var body model.MenuItemDetail
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.UpdateMenuItemDetailParams{
		ID:         req.ID,
		ItemID:     body.ItemID,
		GroupID:    body.GroupID,
		CategoryID: body.CategoryID,
		SizeID:     body.SizeID,
		Cost:       body.Cost,
		Price:      body.Price,
		OutletID:   body.OutletID,
		PrinterID:  body.PrinterID,
		VatID:      body.VatID,
		TerminalID: body.TerminalID,
	}

	err := server.store.UpdateMenuItemDetail(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.UPDATE_SUCCESS, ctx)
}

func (server *Server) DeleteMenuItemDetial(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	err := server.store.DeteleMenuItemDetail(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.DELETE_SUCCESS, ctx)

}
