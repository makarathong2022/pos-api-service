package api

import (
	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/model/request"
	"github.com/boincompany/pos_api_service/model/response"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) getMenuModifies(ctx *gin.Context) {
	var req request.PageInfo

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.GetMenuModifiesParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	modifies, err := server.store.GetMenuModifies(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	total := len(modifies)
	response.OkWithDetailed(
		response.PageResult{
			Result:   modifies,
			Total:    total,
			Page:     req.PageID,
			PageSize: req.PageSize,
			HasNext:  total == int(req.PageSize),
		}, utils.GET_SUCCESS, ctx)
}

func (server *Server) getMenuModify(ctx *gin.Context) {
	var req request.GetById
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	modify, err := server.store.GetMenuModify(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithData(modify, ctx)
}

func (server *Server) createMenuModify(ctx *gin.Context) {
	var body model.MenuModify

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.CreateMenuModifyParams{
		ModifyCd:    body.ModifyCd,
		ModifyName:  body.ModifyName,
		Sort:        body.Sort,
		Description: body.Description,
	}

	modify, err := server.store.CreateMenuModify(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithData(modify, ctx)
}

func (server *Server) updateMenuModify(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	var body model.MenuModify

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.UpdateMenuModifyParams{
		ID:          req.ID,
		ModifyCd:    body.ModifyCd,
		ModifyName:  body.ModifyName,
		Description: body.Description,
		Sort:        body.Sort,
	}

	err := server.store.UpdateMenuModify(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.UPDATE_SUCCESS, ctx)
}

func (server *Server) deleteMenuModify(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	err := server.store.DeleteMenuModify(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.FailWithMessage(utils.DELETE_SUCCESS, ctx)

}
