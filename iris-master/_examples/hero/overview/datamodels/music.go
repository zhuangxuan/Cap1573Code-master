package datamodels

type Music struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	URL           string `json:"url"`
	Poster        string `json:"poster"`
	Genre         string `json:"genre"`
	Album         string `json:"album"`
	Artist        string `json:"artist"`
	Year          int    `json:"year"`
	Lyrics        string `json:"lyrics"`
	LyricsURL     string `json:"lyrics_url"`
	LyricsType    string `json:"lyrics_type"`
	LyricsURLType string `json:"lyrics_url_type"`
}
