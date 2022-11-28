package api

import (
	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/model/request"
	"github.com/boincompany/pos_api_service/model/response"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) GetMenuItemModifies(ctx *gin.Context) {
	var req request.PageInfo

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.GetMenuItemModifiesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	menuItemModifies, err := server.store.GetMenuItemModifies(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	total := len(menuItemModifies)

	response.OkWithDetailed(response.PageResult{
		Total:    total,
		Result:   menuItemModifies,
		Page:     req.PageID,
		PageSize: req.PageSize,
		HasNext:  total == int(req.PageSize),
	}, utils.GET_SUCCESS, ctx)

}

func (server *Server) GetMenuItemModify(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	menuItemModify, err := server.store.GetMenuItemModify(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithData(menuItemModify, ctx)
}

func (server *Server) CreateMenuItemModify(ctx *gin.Context) {
	var body model.MenuItemModify

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.CreateMenuItemModifyParams{
		ItemCd:   body.ItemCd,
		ModifyCd: body.ModifyCd,
	}

	menuItemModify, err := server.store.CreateMenuItemModify(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithData(menuItemModify, ctx)

}

func (server *Server) UpdateMenuItemModify(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	var body model.MenuItemModify

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.UpdateMenuItemModifyParams{
		ItemCd:   body.ItemCd,
		ModifyCd: body.ModifyCd,
		ID:       req.ID,
	}

	err := server.store.UpdateMenuItemModify(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.UPDATE_SUCCESS, ctx)
}
