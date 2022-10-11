package api

import (
	"net/http"
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/gin-gonic/gin"
)

func (server *Server) getMenuGroups(ctx *gin.Context) {
	var req model.Body
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetMenuGroupsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	groups, err := server.store.GetMenuGroups(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	req.Result = groups
	req.HasNext = len(groups) == int(req.PageSize)

	ctx.JSON(http.StatusOK, req)
}

func (server *Server) getMenuGroup(ctx *gin.Context) {
	var req model.Params

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	group, err := server.store.GetMenuGroup(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, group)

}

func (server *Server) createMenuGroup(ctx *gin.Context) {
	var req model.MenuGroup

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMenuGroupParams{
		GroupCd:     req.GroupCd,
		GroupName:   req.GroupName,
		Sort:        req.Sort,
		Description: req.Description,
	}

	group, err := server.store.CreateMenuGroup(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, group)
}

func (server *Server) updateMenuGroup(ctx *gin.Context) {
	var param model.Params

	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req model.MenuGroup

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateMenuGroupParams{
		ID:          param.ID,
		GroupCd:     req.GroupCd,
		GroupName:   req.GroupName,
		Sort:        req.Sort,
		Description: req.Description,
	}
	err := server.store.UpdateMenuGroup(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "One record updated successfully")
}

func (server *Server) deleteMenuGroup(ctx *gin.Context) {
	var req model.Params
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteMenuGroupParams{
		ID:        req.ID,
		DeletedAt: time.Now(),
	}

	err := server.store.DeleteMenuGroup(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "One record delete successfully")
}
