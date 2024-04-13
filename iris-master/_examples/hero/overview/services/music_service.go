package services

import (
	"fmt"

	"github.com/kataras/iris/_examples/hero/overview/repositories"
)

type MusicService interface {
	ShowMusicName() string
}

type MusicServiceManager struct {
	repo repositories.MusicRepository
}

func NewMusicServiceManager(repo repositories.MusicRepository) MusicService {
	return &MusicServiceManager{repo: repo}
}

func (m MusicServiceManager) ShowMusicName() string {
	fmt.Println("我们获取到的音乐名称为：" + m.repo.GetMusicName())
	return "我们获取到的音乐名称为：" + m.repo.GetMusicName()
}
