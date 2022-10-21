package api

import (
	"net/http"
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/model/request"
	"github.com/boincompany/pos_api_service/model/response"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) getMenuSizes(ctx *gin.Context) {

	var req request.PageInfo

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.GetMenuSizesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	sizes, err := server.store.GetMenuSizes(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	total := len(sizes)

	response.OkWithDetailed(
		response.PageResult{
			Result:   sizes,
			Total:    total,
			Page:     req.PageID,
			PageSize: req.PageSize,
			HasNext:  total == int(req.PageSize),
		}, utils.GET_SUCCESS, ctx)

}

func (server *Server) getMenuSize(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
	}

	size, err := server.store.GetMenuSize(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithData(size, ctx)
}

func (server *Server) createMenuSize(ctx *gin.Context) {
	var req model.MenuSize

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	arg := db.CreateMenuSizeParams{
		SizeCd:      req.SizeCd,
		SizeName:    req.SizeName,
		Sort:        req.Sort,
		Description: req.Description,
	}

	size, err := server.store.CreateMenuSize(ctx, arg)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithData(size, ctx)
}

func (server *Server) updateMenuSize(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	var body model.MenuSize

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateMenuSizeParams{
		ID:          req.ID,
		SizeCd:      body.SizeCd,
		SizeName:    body.SizeName,
		Sort:        body.Sort,
		Description: body.Description,
	}

	err := server.store.UpdateMenuSize(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}
	response.OkWithMessage(utils.UPDATE_SUCCESS, ctx)
}

func (server *Server) deleteMenuSize(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.DeleteMenuSizeParams{
		ID:        req.ID,
		DeletedAt: time.Now(),
	}

	err := server.store.DeleteMenuSize(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}
	response.OkWithMessage(utils.DELETE_SUCCESS, ctx)

}
