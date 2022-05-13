package api

import (
	db "github.com/Abdelrhman-Hosny/go_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up the routes
func NewServer(store *db.Store) *Server {

	server := &Server{store: store}
	router := gin.Default()

	// account routes
	router.POST("/accounts", server.createAccount)

	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.getAccounts)

	server.router = router

	return server
}

func (server *Server) Run(addr string) error {
	return server.router.Run(addr)
}

func errorRespone(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
