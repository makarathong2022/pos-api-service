package api

import (
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/model/request"
	"github.com/boincompany/pos_api_service/model/response"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (server *Server) getTerminals(ctx *gin.Context) {
	var req request.PageInfo

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.GetTerminalsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	terminals, err := server.store.GetTerminals(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	total := len(terminals)

	response.OkWithDetailed(response.PageResult{
		Total:    total,
		HasNext:  total == int(req.PageSize),
		Result:   terminals,
		Page:     req.PageID,
		PageSize: req.PageSize,
	}, utils.GET_SUCCESS, ctx)

}

func (server *Server) getTerminal(ctx *gin.Context) {
	var req request.GetById
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	terminal, err := server.store.GetTerminal(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithDetailed(terminal, utils.GET_SUCCESS, ctx)
}

func (server *Server) createNewTerminal(ctx *gin.Context) {
	var body model.Terminal

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.CreateTerminalParams{
		TerminalCd:   body.TerminalCd,
		TerminalName: body.TerminalName,
		Sort:         body.Sort,
		IpAddress:    body.IpAddress,
		Description:  body.Description,
	}

	terminal, err := server.store.CreateTerminal(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithDetailed(terminal, utils.CREATE_SUCCESS, ctx)

}

func (server *Server) updateTerminal(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	var body model.Terminal

	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.UpdateTerminalParams{
		ID:           req.ID,
		TerminalCd:   body.TerminalCd,
		TerminalName: body.TerminalName,
		Sort:         body.Sort,
		IpAddress:    body.IpAddress,
		UpdatedAt:    time.Now(),
		Description:  body.Description,
	}

	err := server.store.UpdateTerminal(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.UPDATE_SUCCESS, ctx)

}

func (server *Server) deleteTerminal(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.DeleteTerminalParams{
		ID:        req.ID,
		DeletedAt: time.Now(),
	}

	err := server.store.DeleteTerminal(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.DELETE_SUCCESS, ctx)
}
