package api

import (
	db "bank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server for HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// New server creates the routing and server exec.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// middlewares
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// accounts
	router.POST("/api/accounts", server.createAccount)
	router.GET("/api/accounts/:id", server.getAccount)
	router.GET("/api/accounts", server.listAccounts)
	// transfers
	router.POST("/api/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start runs the server on a port
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
