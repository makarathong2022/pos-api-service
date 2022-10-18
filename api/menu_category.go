package api

import (
	"net/http"
	"time"
	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/gin-gonic/gin"
)

func (server *Server) getMenuCategories(ctx *gin.Context) {
	var req model.Body

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetMenuCategoriesParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	categories, err := server.store.GetMenuCategories(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	req.Result = categories
	req.HasNext = len(categories) == int(req.PageSize)
	req.Total = len(categories)
	ctx.JSON(http.StatusOK, req)

}

func (server *Server) getMenuCategory(ctx *gin.Context) {
	var req model.Params

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.GetMenuCategory(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)

}

func (server *Server) createMenuCategory(ctx *gin.Context) {
	var req model.Category

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (server *Server) updateMenuCategory(ctx *gin.Context) {
	var req model.Category
	var param model.Params

	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateMenuCategoryParams{
		ID:           param.ID,
		CategoryCd:   req.CategoryCd,
		CategoryName: req.CategoryName,
		Sort:         req.Sort,
		Description:  req.Description,
		UpdatedAt:    time.Now(),
	}

	err := server.store.UpdateMenuCategory(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusBadRequest, "One record update successfully")
}

func (server *Server) deleteMenuCategory(ctx *gin.Context) {
	var req model.Params

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteMenuCategory(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "One record has been deleted")
}
