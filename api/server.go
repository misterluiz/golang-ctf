package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/misterluiz/golang-ctf/db/sqlc"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//colocar as rotas aqui
	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	router.GET("/user/:id", server.getUserById)
	server.router = router

	return server
}

func (server *Server) Star(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has error:": err.Error()}
}
