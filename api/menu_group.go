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

func (server *Server) GetMenuGroups(ctx *gin.Context) {
	var req request.PageInfo
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.GetMenuGroupsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	groups, err := server.store.GetMenuGroups(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	total := len(groups)

	response.OkWithDetailed(response.PageResult{
		Total:    total,
		HasNext:  total == int(req.PageSize),
		Result:   groups,
		Page:     req.PageID,
		PageSize: req.PageSize,
	}, utils.GET_SUCCESS, ctx)
}

func (server *Server) GetMenuGroup(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	group, err := server.store.GetMenuGroup(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithDetailed(group, utils.CREATE_SUCCESS, ctx)
}

func (server *Server) CreateMenuGroup(ctx *gin.Context) {
	var req model.MenuGroup

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
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
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithDetailed(group, utils.CREATE_SUCCESS, ctx)
}

func (server *Server) UpdateMenuGroup(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	var body model.MenuGroup

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.UpdateMenuGroupParams{
		ID:          req.ID,
		GroupCd:     body.GroupCd,
		GroupName:   body.GroupName,
		Sort:        body.Sort,
		Description: body.Description,
	}
	err := server.store.UpdateMenuGroup(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.UPDATE_SUCCESS, ctx)
}

func (server *Server) DeleteMenuGroup(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.DeleteMenuGroupParams{
		ID:        req.ID,
		DeletedAt: time.Now(),
	}

	err := server.store.DeleteMenuGroup(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.DELETE_SUCCESS, ctx)
}
