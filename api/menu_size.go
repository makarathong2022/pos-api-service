package api

import (
	"net/http"
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/model/response"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

// GetExaMenuSizes
// @Tags      ExaMenuSizes
// @Summary   分页获取权限客户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}
// @Router    /sizes [get]
func (server *Server) getMenuSizes(ctx *gin.Context) {
	var req model.Body

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

	response.OkWithDetailed(response.PageResult{
		List:     sizes,
		Total:    int64(total),
		Page:     int(req.PageID),
		PageSize: int(req.PageSize),
		HasNext:  total == int(req.PageSize),
	}, utils.GET_SUCCESS, ctx)

}

func (server *Server) getMenuSize(ctx *gin.Context) {
	var req model.Params

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
	var params model.Params

	if err := ctx.ShouldBindUri(&params); err != nil {
		response.FailWithMessage(errRes(err), ctx)
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
		response.FailWithMessage(errRes(err), ctx)
		return
	}
	response.OkWithMessage("One data update successfully", ctx)
}

func (server *Server) deleteMenuSize(ctx *gin.Context) {
	var params model.Params

	if err := ctx.ShouldBindUri(params); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.DeleteMenuSizeParams{
		ID:        params.ID,
		DeletedAt: time.Now(),
	}

	err := server.store.DeleteMenuSize(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}
	response.OkWithMessage("One record deleted successfully", ctx)

}
