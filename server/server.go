package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

func (server *Server) Run() {
	router := gin.Default()

	router.GET("/announce", AnnounceHandler(server))

	http.ListenAndServe(":5010", router)
}
