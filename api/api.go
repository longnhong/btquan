package api

import (
	"btquan/api/auth"
	"btquan/system"
	"github.com/gin-gonic/gin"
)

func InitApi(root *gin.RouterGroup, tkWorker *system.VideoWorker) {
	auth.NewAuthenServer(root, "auth")
}
