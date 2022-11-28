package api

import (
	"fmt"

	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/token"
	"github.com/boincompany/pos_api_service/utils"
	"github.com/gin-gonic/gin"
)

// Server HTTP requests for our banking service
type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config utils.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{store: store, tokenMaker: tokenMaker, config: config}

	/*
		- After create Gin Router we call binding.Validator.Engine() to get current validator engin that gin is using
		and note that this function will return a general interface type. Which by default, is a pointer to the validator object
		of the go-playground validator package
		- *&validator.Validate{}: convert the output to a validator. Validator pointer
		- if condition "ok" then we can call v.RegisterValidation
	*/
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("currency", validCurrency)
	// }
	server.setupRouter()
	return server, nil
}
func (server *Server) setupRouter() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})
	router.GET("/", server.showDefaultRoute)

	router.POST("/items", server.CreateMenuItem)
	router.GET("/items", server.GetMenuItems)
	router.GET("/items/:id", server.GetMenuItem)
	router.PUT("/items/", server.UpdateMenuItem)
	router.DELETE("/items/", server.DeleteMenuItem)

	router.GET("/categories", server.GetMenuCategories)
	router.GET("/categories/:id", server.GetMenuCategory)
	router.POST("/categories", server.CreateMenuCategory)
	router.PUT("/categories/:id", server.UpdateMenuCategory)
	router.DELETE("/categories/:id", server.DeleteMenuCategory)

	router.GET("/groups", server.GetMenuGroups)
	router.GET("/groups/:id", server.GetMenuGroup)
	router.POST("/groups", server.CreateMenuGroup)
	router.PUT("/groups/:id", server.UpdateMenuGroup)
	router.DELETE("/groups/:id", server.DeleteMenuGroup)

	router.GET("/modifies", server.GetMenuModifies)
	router.GET("/modifies/:id", server.GetMenuModify)
	router.POST("/modifies", server.CreateMenuModify)
	router.PUT("/modifies", server.UpdateMenuModify)
	router.DELETE("/modifies", server.DeleteMenuModify)

	router.GET("/sizes", server.GetMenuSizes)
	router.GET("/sizes/:id", server.GetMenuSize)
	router.POST("/sizes", server.CreateMenuSize)
	router.PUT("/sizes/:id", server.UpdateMenuSize)
	router.DELETE("/sizes/:id", server.DeleteMenuSize)

	router.GET("/outlets", server.GetOutlets)
	router.GET("/outlets/:id", server.GetOutlet)
	router.POST("/outlets", server.CreateNewOutlet)
	router.PUT("/outlets/:id", server.UpdateOutlet)
	router.DELETE("/outlets/:id", server.DeleteOutlet)

	router.GET("/printers", server.getPrinters)
	router.GET("/printers/:id", server.getPrinter)
	router.POST("/printers", server.createNewPrinter)
	router.PUT("/printers/:id", server.updatePrinter)
	router.DELETE("/printers/:id", server.deletePrinter)

	router.GET("/terminals", server.GetTerminals)
	router.GET("/terminals/:id", server.GetTerminal)
	router.POST("/terminals", server.CreateNewTerminal)
	router.PUT("/terminals/:id", server.UpdateTerminal)
	router.DELETE("/terminals/:id", server.DeleteTerminal)

	router.GET("/item_modifies", server.GetMenuItemModifies)
	router.GET("/item_modifies/:id", server.GetMenuItemModify)
	router.POST("/item_modifies", server.CreateMenuItemModify)
	router.PUT("/item_modifies/:id", server.UpdateMenuItemModify)

	router.GET("/menu-item-details", server.GetMenuItemDetials)
	router.GET("/menu-item-details/:id", server.GetMenuItemDetial)
	router.POST("/menu-item-details", server.CreateMenuItemDetail)
	router.PUT("/menu-item-details/:id", server.UpdateMenuItemDetial)
	router.DELETE("/menu-item-details/:id", server.DeleteMenuItemDetial)

	router.GET("/taxs", server.GetTaxs)
	router.GET("/taxs/:id", server.GetTax)
	router.POST("/taxs", server.CreateTax)
	router.PUT("/taxs", server.UpdateTax)
	router.DELETE("/taxs", server.DeleteTax)
	// router.POST("/users/login", server.loginUser)

	// "/" This slash is the path prefix of all routes in this group
	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.listAccounts)
	// authRoutes.POST("/transfers", server.createTransfer)
	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errRes(err error) string {
	return err.Error()
}
