package api

import (
	db "bank/db/sqlc"

	"github.com/gin-gonic/gin"
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

	router.POST("/api/accounts", server.createAccount)
	router.GET("/api/accounts/:id", server.getAccount)
	router.GET("/api/accounts", server.listAccounts)

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
