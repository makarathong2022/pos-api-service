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

func (server *Server) getMenuCategories(ctx *gin.Context) {
	var req request.PageInfo

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.GetMenuCategoriesParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	categories, err := server.store.GetMenuCategories(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	total := len(categories)

	response.OkWithDetailed(response.PageResult{
		PageSize: req.PageSize,
		Page:     req.PageID,
		Result:   categories,
		Total:    total,
		HasNext:  total == int(req.PageSize),
	}, utils.GET_SUCCESS, ctx)

}

func (server *Server) getMenuCategory(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	category, err := server.store.GetMenuCategory(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithData(category, ctx)
}

func (server *Server) createMenuCategory(ctx *gin.Context) {
	var req model.Category

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.CreateMenuCategoryParams{
		CategoryCd:   req.CategoryCd,
		CategoryName: req.CategoryName,
		Sort:         req.Sort,
		Description:  req.Description,
	}

	category, err := server.store.CreateMenuCategory(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithData(category, ctx)
}

func (server *Server) updateMenuCategory(ctx *gin.Context) {
	var body model.Category
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.UpdateMenuCategoryParams{
		ID:           req.ID,
		CategoryCd:   body.CategoryCd,
		CategoryName: body.CategoryName,
		Sort:         body.Sort,
		Description:  body.Description,
		UpdatedAt:    time.Now(),
	}

	err := server.store.UpdateMenuCategory(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.UPDATE_SUCCESS, ctx)
}

func (server *Server) deleteMenuCategory(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	err := server.store.DeleteMenuCategory(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.DELETE_SUCCESS, ctx)
}
