package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/model"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) createMenuItem(ctx *gin.Context) {
	var req model.MenuItem

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}
	ctx.JSON(http.StatusOK, menuItem)
}

func (server *Server) getMenuItems(ctx *gin.Context) {
	var req model.Body
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetMenuItemsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	fmt.Println(arg)

	menuItems, err := server.store.GetMenuItems(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	req.HasNext = len(menuItems) == int(req.PageSize)
	req.Total = len(menuItems)
	req.Result = menuItems
	ctx.JSON(http.StatusOK, req)
}

func (server *Server) getMenuItem(ctx *gin.Context) {
	var res model.MenuItemResponse
	if err := ctx.ShouldBindUri(&res); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	menuItem, err := server.store.GetMenuItem(ctx, int64(res.ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, menuItem)

}

func (server *Server) updateMenuItem(ctx *gin.Context) {

	var params model.Params
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req model.MenuItem
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateMenuItemParams{
		ID:          int64(params.ID),
		ItemCd:      req.ItemCd,
		ItemName:    req.ItemName,
		Abv:         req.Abv,
		Description: req.Description,
		OutletID:    req.OutletID,
		UpdatedAt:   time.Now(),
	}

	err := server.store.UpdateMenuItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, "One Record has been updated")
}

// Delete Menu Item
func (server *Server) deleteMenuItem(ctx *gin.Context) {
	var req model.MenuItemResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req.ID)

	err := server.store.DeleteMenuItem(ctx, int64(req.ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "One Record has been deleted")
}
