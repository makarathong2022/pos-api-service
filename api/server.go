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
	router.POST("/items", server.createMenuItem)
	router.GET("/items", server.getMenuItems)
	router.GET("/items/:id", server.getMenuItem)
	router.PUT("/items/", server.updateMenuItem)
	router.DELETE("/items/", server.deleteMenuItem)

	router.GET("/categories", server.getMenuCategories)
	router.GET("/categories/:id", server.getMenuCategory)
	router.POST("/categories", server.createMenuCategory)
	router.PUT("/categories/:id", server.updateMenuCategory)
	router.DELETE("/categories/:id", server.deleteMenuCategory)

	router.GET("/groups", server.getMenuGroups)
	router.GET("/groups/:id", server.getMenuGroup)
	router.POST("/groups", server.createMenuGroup)
	router.PUT("/groups/:id", server.updateMenuGroup)
	router.DELETE("/groups/:id", server.deleteMenuGroup)

	router.GET("/modifies", server.getMenuModifies)
	router.GET("/modifies/:id", server.getMenuModify)
	router.POST("/modifies", server.createMenuModify)
	router.PUT("/modifies", server.updateMenuModify)
	router.DELETE("/modifies", server.deleteMenuModify)

	router.GET("/sizes", server.getMenuSizes)
	router.GET("/sizes/:id", server.getMenuSize)
	router.POST("/sizes", server.createMenuSize)
	router.PUT("/sizes/:id", server.updateMenuSize)
	router.DELETE("/sizes/:id", server.deleteMenuSize)

	router.GET("/outlets", server.getOutlets)
	router.GET("/outlets/:id", server.getOutlet)
	router.POST("/outlets", server.createNewOutlet)
	router.PUT("/outlets/:id", server.updateOutlet)
	router.DELETE("/outlets/:id", server.deleteOutlet)

	router.GET("/printers", server.getPrinters)
	router.GET("/printers/:id", server.getPrinter)
	router.POST("/printers", server.createNewPrinter)
	router.PUT("/printers/:id", server.updatePrinter)
	router.DELETE("/printers/:id", server.deletePrinter)

	router.GET("/terminals", server.getTerminals)
	router.GET("/terminals/:id", server.getTerminal)
	router.POST("/terminals", server.createNewTerminal)
	router.PUT("/terminals/:id", server.updateTerminal)
	router.DELETE("/terminals/:id", server.deleteTerminal)

	router.POST("/menu-item-details", server.createMenuItemDetail)

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

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func errRes(err error) string {
	return err.Error()
}
