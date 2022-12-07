package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "pixel-thc-backend-go/db/sqlc"
	"pixel-thc-backend-go/token"
	"pixel-thc-backend-go/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// unauthenticated routes
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	// authenticated routes
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// accounts
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)

	// transfers
	authRoutes.POST("/transfers", server.createTransfer)

	// product categories
	authRoutes.GET("/product-categories", server.listProductCategories)
	authRoutes.GET("/product-categories/:id", server.getProductCategory)
	authRoutes.POST("/product-categories", server.createProductCategory)

	// uoms
	authRoutes.GET("/uoms", server.listUoms)
	authRoutes.GET("/uoms/:id", server.getUom)
	authRoutes.POST("/uoms", server.createUom)

	// strains
	authRoutes.GET("/strains", server.listStrains)
	authRoutes.GET("/strains/:id", server.getStrain)
	authRoutes.POST("/strains", server.createStrain)

	// item types
	authRoutes.GET("/item-types", server.listItemTypes)
	authRoutes.GET("/item-types/:id", server.getItemType)
	authRoutes.POST("/item-types", server.createItemType)

	// items
	authRoutes.GET("/items", server.listItems)
	authRoutes.GET("/items/:id", server.getItem)
	authRoutes.POST("/items", server.createItem)

	// lab tests
	authRoutes.GET("/lab-tests", server.listLabTests)
	authRoutes.GET("/lab-tests/:id", server.getLabTest)
	authRoutes.POST("/lab-tests", server.createLabTest)

	// package tags
	authRoutes.GET("/package-tags", server.listPackageTags)
	authRoutes.GET("/package-tags/:id", server.getPackageTag)
	authRoutes.GET("/package-tags/tag/:tag-number", server.getPackageTagByTagNumber)
	authRoutes.POST("/package-tags", server.createPackageTag)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
