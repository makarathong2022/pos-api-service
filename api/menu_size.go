package api

import (
	"net/http"
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/gin-gonic/gin"
)

func (server *Server) getMenuSizes(ctx *gin.Context) {
	var req model.Body

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetMenuSizesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	sizes, err := server.store.GetMenuSizes(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	countSize := len(sizes)
	req.HasNext = countSize == int(req.PageSize)
	req.Total = countSize
	req.Result = sizes
	ctx.JSON(http.StatusOK, req)

}

func (server *Server) getMenuSize(ctx *gin.Context) {
	var req model.Params

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	size, err := server.store.GetMenuSize(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, size)
}

func (server *Server) createMenuSize(ctx *gin.Context) {
	var req model.MenuSize

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, size)
}

func (server *Server) updateMenuSize(ctx *gin.Context) {
	var params model.Params

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req model.MenuSize

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateMenuSizeParams{
		ID:          params.ID,
		SizeCd:      req.SizeCd,
		SizeName:    req.SizeName,
		Sort:        req.Sort,
		Description: req.Description,
	}

	err := server.store.UpdateMenuSize(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Update successfully")
}

func (server *Server) deleteMenuSize(ctx *gin.Context) {
	var params model.Params

	if err := ctx.ShouldBindUri(params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteMenuSizeParams{
		ID:        params.ID,
		DeletedAt: time.Now(),
	}

	err := server.store.DeleteMenuSize(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "One record deleted successfully")

}
