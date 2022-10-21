package api

import (
	"log"
	"net/http"
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/boincompany/pos_api_service/model/request"
	"github.com/boincompany/pos_api_service/model/response"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) createMenuItem(ctx *gin.Context) {
	var req model.MenuItem

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.CreateMenuItemParams{
		ItemCd:      req.ItemCd,
		ItemName:    req.ItemName,
		Abv:         req.Abv,
		Sort:        req.Sort,
		Description: req.Description,
		OutletID:    req.OutletID,
	}

	menuItem, err := server.store.CreateMenuItem(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violoation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithDetailed(menuItem, utils.CREATE_SUCCESS, ctx)
}

func (server *Server) getMenuItems(ctx *gin.Context) {
	var req request.PageInfo

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.GetMenuItemsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	menuItems, err := server.store.GetMenuItems(ctx, arg)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	total := len(menuItems)

	response.OkWithDetailed(response.PageResult{
		Total:    total,
		HasNext:  total == int(req.PageSize),
		Result:   menuItems,
		Page:     req.PageID,
		PageSize: req.PageSize,
	}, utils.GET_SUCCESS, ctx)

}

func (server *Server) getMenuItem(ctx *gin.Context) {
	var res request.GetById
	if err := ctx.ShouldBindUri(&res); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	menuItem, err := server.store.GetMenuItem(ctx, int64(res.ID))

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithDetailed(menuItem, utils.GET_SUCCESS, ctx)

}

func (server *Server) updateMenuItem(ctx *gin.Context) {

	var req request.GetById
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	var body model.MenuItem
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	arg := db.UpdateMenuItemParams{
		ID:          req.ID,
		ItemCd:      body.ItemCd,
		ItemName:    body.ItemName,
		Abv:         body.Abv,
		Description: body.Description,
		OutletID:    body.OutletID,
		UpdatedAt:   time.Now(),
	}

	err := server.store.UpdateMenuItem(ctx, arg)
	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.UPDATE_SUCCESS, ctx)
}

func (server *Server) deleteMenuItem(ctx *gin.Context) {
	var req request.GetById

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	err := server.store.DeleteMenuItem(ctx, req.ID)

	if err != nil {
		response.FailWithMessage(errRes(err), ctx)
		return
	}

	response.OkWithMessage(utils.DELETE_SUCCESS, ctx)
}
