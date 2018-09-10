package main

import (
	// 1. init first
	_ "LongTM/btq/btquan/init"
	"net/http"
	// 2. iniit 2nd
	"LongTM/btq/btquan/api"
	"LongTM/btq/btquan/common"
	"LongTM/btq/btquan/middleware"
	"LongTM/btq/btquan/room"

	"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	router := gin.New()
	router.StaticFS("/admin", http.Dir("./admin")).Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/x-javascript")
		c.Next()
	})
	router.Use(middleware.AddHeader(), gin.Logger(), middleware.Recovery())
	//static
	// router.StaticFS("/static", http.Dir("./upload"))
	// router.StaticFS("/app", http.Dir("./app"))

	//api
	rootAPI := router.Group("/api")
	api.InitApi(rootAPI)
	//ws
	room.NewRoomServer(router.Group("/room"))
	router.Run(common.ConfigSystemBooking.PortBooking)
}
