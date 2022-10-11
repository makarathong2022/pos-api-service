package api

import (
	"net/http"
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	utils "github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) getOutlets(ctx *gin.Context) {
	var req model.Body

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetOutletsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	outlets, err := server.store.GetOutlets(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	req.HasNext = len(outlets) == int(req.PageSize)
	req.Result = outlets
	req.Total = len(outlets)
	ctx.JSON(http.StatusOK, req)
}

func (server *Server) getOutlet(ctx *gin.Context) {
	var params model.Params

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	outlet, err := server.store.GetOutlet(ctx, params.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, outlet)
}

func (server *Server) createNewOutlet(ctx *gin.Context) {
	var req model.Outlet

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, outlet)
}

func (server *Server) updateOutlet(ctx *gin.Context) {
	var params model.Params

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req model.Outlet

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdateOutletParams{
		ID:          params.ID,
		OutletCd:    req.OutletCd,
		OutletName:  req.OutletName,
		IpAddress:   req.IpAddress,
		Description: req.Description,
		UpdatedAt:   time.Now(),
	}

	err := server.store.UpdateOutlet(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.UPDATE_SUCCESS)
}

func (server *Server) deleteOutlet(ctx *gin.Context) {
	var params model.Params

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteOutletParams{
		ID:        params.ID,
		DeletedAt: time.Now(),
	}

	err := server.store.DeleteOutlet(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.DELETE_SUCCESS)
}
