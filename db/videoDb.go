package db

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Video struct {
	ID        string `gorm:"primarykey"`
	CreateTime string
	VideoTitle string
	VideoFileName string
	VideoSeconds float32
	VideoWidth int
	VideoHeight int
	LikeNum int
	Status int
	CommentNum int
}

func AddVideo(video *Video) error {
	r := DB.Create(video)
	if r.Error != nil {
		if mysqlErr, ok := r.Error.(*mysql.MySQLError); ok{
			//id重复
			if mysqlErr.Number == 1062 {
				return mysqlErr
			}
		}
		//未知错误
		panic(r.Error)
	}
	return nil
}

func GetVideoById(id string) (*Video, error) {
	video := Video{}
	r := DB.First(&video, "id = ?", id)
	if r.Error != nil {
		//video 不存在
		if r.Error == gorm.ErrRecordNotFound{
			return nil, r.Error
		}
		//未知错误
		panic(r.Error)
	}
	return &video, nil
}

//返回从sinceId开始的num个video
func GetVideosBySinceId(sinceId string, num int) ([]*Video, error) {
	videos := make([]*Video, num)
	r := DB.Where("id > ?", sinceId).Limit(num).Find(&videos)
	if r.Error != nil {
		//未知错误
		panic(r.Error)
	}
	getVideoNum := 0
	for _, v := range videos {
		if v != nil {
			getVideoNum ++
		}else {
			break
		}
	}
	return videos[:getVideoNum], nil
}
