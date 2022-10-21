package api

import (
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/model/request"
	"github.com/boincompany/pos_api_service/model/response"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) getPrinters(ctx *gin.Context) {
	var req request.PageInfo

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.GetPrintersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	printers, err := server.store.GetPrinters(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	total := len(printers)

	response.OkWithDetailed(response.PageResult{
		Total:    total,
		HasNext:  total == int(req.PageSize),
		Result:   printers,
		Page:     req.PageID,
		PageSize: req.PageSize,
	}, utils.GET_SUCCESS, ctx)
}

func (server *Server) getPrinter(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	printer, err := server.store.GetPrinter(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithDetailed(printer, utils.GET_SUCCESS, ctx)
}

func (server *Server) createNewPrinter(ctx *gin.Context) {
	var body model.Printer

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.CreatePrinterParams{
		PrintCd:     body.PrinterCd,
		PrintName:   body.PrinterName,
		Sort:        body.Sort,
		Description: body.Description,
		IpAddress:   body.IpAddress,
	}

	printer, err := server.store.CreatePrinter(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithDetailed(printer, utils.CREATE_SUCCESS, ctx)

}

func (server *Server) updatePrinter(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	var body model.Printer

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.UpdatePrinterParams{
		ID:          req.ID,
		PrintCd:     body.PrinterCd,
		PrintName:   body.PrinterName,
		IpAddress:   body.IpAddress,
		Sort:        body.Sort,
		Description: body.Description,
	}

	err := server.store.UpdatePrinter(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.UPDATE_SUCCESS, ctx)

}

func (server *Server) deletePrinter(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.DeletePrinterParams{
		ID:        req.ID,
		DeletedAt: time.Now(),
	}
	err := server.store.DeletePrinter(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.DELETE_SUCCESS, ctx)

}
