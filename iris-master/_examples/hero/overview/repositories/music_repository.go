package repositories

import "github.com/kataras/iris/_examples/hero/overview/datamodels"

type MusicRepository interface {
	GetMusicName() string
}

type MusicManager struct{}

func NewMusicManager() MusicRepository {
	return &MusicManager{}
}

func (m *MusicManager) GetMusicName() string {
	music := &datamodels.Music{Name: "music"}
	return music.Name
}
