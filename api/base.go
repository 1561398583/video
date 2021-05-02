package api

import (
	"github.com/gin-gonic/gin"
	"yx.com/videos/utils"
)

var Logger *utils.PdLog

func RegistApi(r *gin.Engine)  {
	//user api
	r.GET("/login", GetLoginPage)
	r.POST("/login", Login)

	//video api
	r.GET("/getVideosInfor", GetVideosInfo)
}
