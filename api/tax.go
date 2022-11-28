package api

import (
	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/model/request"
	"github.com/boincompany/pos_api_service/model/response"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) GetTaxs(ctx *gin.Context) {
	var req request.PageInfo

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.GetTaxsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	taxs, err := server.store.GetTaxs(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	total := len(taxs)

	response.OkWithDetailed(response.PageResult{
		Result:   taxs,
		Total:    total,
		HasNext:  total == int(req.PageSize),
		Page:     req.PageID,
		PageSize: req.PageSize,
	}, utils.GET_SUCCESS, ctx)
}

func (server *Server) GetTax(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	tax, err := server.store.GetTax(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithData(tax, ctx)
}

func (server *Server) CreateTax(ctx *gin.Context) {
	var body model.Tax

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.CreateTaxParams{
		VatCd:       body.VatCd,
		VatKey:      body.VatKey,
		VatName:     body.VatName,
		Sort:        body.Sort,
		Description: body.Description,
	}

	tax, err := server.store.CreateTax(ctx, arg)
	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithData(tax, ctx)

}

func (server *Server) UpdateTax(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	var body model.Tax

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.UpdateTaxParams{
		ID:          req.ID,
		VatCd:       body.VatCd,
		VatName:     body.VatName,
		VatKey:      body.VatKey,
		Sort:        body.Sort,
		Description: body.Description,
	}

	err := server.store.UpdateTax(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.UPDATE_SUCCESS, ctx)

}


func (server *Server) DeleteTax(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	err := server.store.DeleteTax(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.DELETE_SUCCESS, ctx)
}
