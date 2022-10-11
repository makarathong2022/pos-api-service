package api

import (
	"net/http"
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) getTerminals(ctx *gin.Context) {
	var req model.Body

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetTerminalsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	terminals, err := server.store.GetTerminals(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	countTerminal := len(terminals)
	req.Result = terminals
	req.Total = countTerminal
	req.HasNext = countTerminal == int(req.PageSize)

	ctx.JSON(http.StatusOK, req)

}

func (server *Server) getTerminal(ctx *gin.Context) {
	var params model.Params
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	terminal, err := server.store.GetTerminal(ctx, params.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, terminal)
}

func (server *Server) createNewTerminal(ctx *gin.Context) {
	var req model.Terminal

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTerminalParams{
		TerminalCd:   req.TerminalCd,
		TerminalName: req.TerminalName,
		Sort:         req.Sort,
		IpAddress:    req.IpAddress,
		Description:  req.Description,
	}

	terminal, err := server.store.CreateTerminal(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, terminal)

}

func (server *Server) updateTerminal(ctx *gin.Context) {
	var params model.Params

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req model.Terminal

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateTerminalParams{
		ID:           params.ID,
		TerminalCd:   req.TerminalCd,
		TerminalName: req.TerminalName,
		Sort:         req.Sort,
		IpAddress:    req.IpAddress,
		UpdatedAt:    time.Now(),
		Description:  req.Description,
	}

	err := server.store.UpdateTerminal(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.UPDATE_SUCCESS)

}

func (server *Server) deleteTerminal(ctx *gin.Context) {
	var params model.Params

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteTerminalParams{
		ID:        params.ID,
		DeletedAt: time.Now(),
	}

	err := server.store.DeleteTerminal(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.DELETE_SUCCESS)
}
