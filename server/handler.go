package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaksonkallio/go-torrent-tracker/torrent"
)

func AnnounceHandler(server *Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		peerId := ctx.Query("peer_id")
		ip := ctx.Query("ip")
		port := ctx.Query("port")
		uploaded := ctx.Query("uploaded")
		downloaded := ctx.Query("downloaded")
		left := ctx.Query("left")
		event := ctx.Query("event")

		announcement, err := torrent.BuildAnnouncement(
			peerId,
			ip,
			port,
			uploaded,
			downloaded,
			left,
			event,
		)

		if err != nil {
			log.Printf("bad announcement: %s", err)
		}

		log.Printf("%#v", announcement)

		ctx.String(http.StatusOK, "hello world!")
	}
}
