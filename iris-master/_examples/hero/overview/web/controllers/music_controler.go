package controllers

import (
	"fmt"

	"github.com/kataras/iris/_examples/hero/overview/repositories"
	"github.com/kataras/iris/_examples/hero/overview/services"
	"github.com/kataras/iris/mvc"
)

type MusicController struct {
}

// 定义controller能处理那些请求
func (c *MusicController) Get() mvc.View {
	musicRepository := repositories.NewMusicManager()
	musicService := services.NewMusicServiceManager(musicRepository)
	musicName := musicService.ShowMusicName()
	fmt.Println("musicName", musicName)
	return mvc.View{
		Name: "/music/index.html",
		Data: musicName,
	}
}
