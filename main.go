package main

import (
	// 1. init first
	_ "btquan/init"
	// 2. iniit 2nd
	"btquan/api"
	"btquan/common"
	"btquan/middleware"
	"btquan/room"
	"btquan/system"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	router := gin.New()
	// router.StaticFS("/admin", http.Dir("./admin")).Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Content-Type", "text/html")
	// 	c.Next()
	// })
	router.Use(middleware.AddHeader(), gin.Logger(), middleware.Recovery())
	//static
	// router.StaticFS("/static", http.Dir("./upload"))
	// router.StaticFS("/app", http.Dir("./app"))
	var tkWorker = system.Start()
	tkWorker.Launch()

	//api
	rootAPI := router.Group("/api")
	api.InitApi(rootAPI, tkWorker)
	//ws
	room.NewRoomServer(router.Group("/room"))
	router.Run(common.ConfigSystemBooking.PortBooking)
}
