package api

import (
	"github.com/gin-gonic/gin"
	"yx.com/videos/ServerConst"
	"yx.com/videos/db"
)

func GetVideosInfo(c *gin.Context){
	sinceVideoId := c.Query("sinceVideoId")
	userId := c.Query("uid")
	videos, err := db.GetVideosBySinceId(sinceVideoId, 5)

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	if len(videos) == 0 {	//说明到达最后一个video了，那么就从头开始
		videos, err = db.GetVideosBySinceId(sinceVideoId, 5)
		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}
	}
	//转换struct
	clientVideos := make([]*VideoInfo2Client, len(videos))
	for i := 0; i < len(videos); i++ {
		v2c := new(VideoInfo2Client)
		v2c.ID = videos[i].ID
		v2c.VideoTitle = videos[i].VideoTitle
		v2c.VideoUrl = ServerConst.HOST + "/" + "assets/videos/" + videos[i].VideoFileName
		v2c.CommentNum = videos[i].CommentNum
		v2c.LikeNum = videos[i].LikeNum
		v2c.VideoSeconds = videos[i].VideoSeconds
		islike := db.IsUserLikeVideo(userId, videos[i].ID)
		v2c.IsLike = islike
		clientVideos[i] = v2c
	}
	c.JSON(200, clientVideos)
}

type VideoInfo2Client struct {
	ID        string `gorm:"primarykey"`
	VideoTitle string
	VideoUrl string
	LikeNum int
	IsLike bool
	CommentNum int
	VideoSeconds float32
}
