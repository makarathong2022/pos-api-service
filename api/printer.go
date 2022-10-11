package api

import (
	"net/http"
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) getPrinters(ctx *gin.Context) {
	var req model.Body

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetPrintersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	printers, err := server.store.GetPrinters(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	req.Result = printers
	req.HasNext = len(printers) == int(req.PageSize)
	req.Total = len(printers)
	ctx.JSON(http.StatusOK, req)

}

func (server *Server) getPrinter(ctx *gin.Context) {
	var params model.Params

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	printer, err := server.store.GetPrinter(ctx, params.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, printer)
}

func (server *Server) createNewPrinter(ctx *gin.Context) {
	var req model.Printer

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePrinterParams{
		PrintCd:     req.PrinterCd,
		PrintName:   req.PrinterName,
		Sort:        req.Sort,
		Description: req.Description,
		IpAddress:   req.IpAddress,
	}

	printer, err := server.store.CreatePrinter(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, printer)

}

func (server *Server) updatePrinter(ctx *gin.Context) {
	var param model.Params

	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req model.Printer

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePrinterParams{
		ID:          param.ID,
		PrintCd:     req.PrinterCd,
		PrintName:   req.PrinterName,
		IpAddress:   req.IpAddress,
		Sort:        req.Sort,
		Description: req.Description,
	}

	err := server.store.UpdatePrinter(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.UPDATE_SUCCESS)

}

func (server *Server) deletePrinter(ctx *gin.Context) {
	var params model.Params

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeletePrinterParams{
		ID:        params.ID,
		DeletedAt: time.Now(),
	}
	err := server.store.DeletePrinter(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.DELETE_SUCCESS)

}
