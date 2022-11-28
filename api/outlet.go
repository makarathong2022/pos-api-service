package api

import (
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/model/request"
	"github.com/boincompany/pos_api_service/model/response"
	utils "github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) GetOutlets(ctx *gin.Context) {
	var req request.PageInfo

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.GetOutletsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	outlets, err := server.store.GetOutlets(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	total := len(outlets)

	response.OkWithDetailed(response.PageResult{
		Total:    total,
		HasNext:  total == int(req.PageSize),
		Result:   outlets,
		Page:     req.PageID,
		PageSize: req.PageSize,
	}, utils.GET_SUCCESS, ctx)

}

func (server *Server) GetOutlet(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	outlet, err := server.store.GetOutlet(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithDetailed(outlet, utils.GET_SUCCESS, ctx)
}

func (server *Server) CreateNewOutlet(ctx *gin.Context) {
	var req model.Outlet

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.CreateOutletParams{
		OutletCd:    req.OutletCd,
		OutletName:  req.OutletName,
		Sort:        req.Sort,
		IpAddress:   req.IpAddress,
		Description: req.Description,
	}

	outlet, err := server.store.CreateOutlet(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithDetailed(outlet, utils.CREATE_SUCCESS, ctx)
}

func (server *Server) UpdateOutlet(ctx *gin.Context) {
	var req request.GetById
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	var body model.Outlet

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.UpdateOutletParams{
		ID:          req.ID,
		OutletCd:    body.OutletCd,
		OutletName:  body.OutletName,
		IpAddress:   body.IpAddress,
		Description: body.Description,
		UpdatedAt:   time.Now(),
	}

	err := server.store.UpdateOutlet(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.UPDATE_SUCCESS, ctx)
}

func (server *Server) DeleteOutlet(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.DeleteOutletParams{
		ID:        req.ID,
		DeletedAt: time.Now(),
	}

	err := server.store.DeleteOutlet(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.DELETE_SUCCESS, ctx)
}
