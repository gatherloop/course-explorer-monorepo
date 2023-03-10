package entity

import "os"

var BaseURL = os.Getenv("BASE_URL")
var Key = os.Getenv("API_KEY")
var ChannelId = os.Getenv("CHANNEL_ID")

type GetResponse struct{
  Items []GetPlaylist `json:"items"`
}

type GetPlaylist struct{
  Snippet GetPlaylistData `json:"snippet"`
}

type GetPlaylistData struct{
  Title string `json:"title"`
  Description string `json:"description"`
  Thumbnails Thumbnails `json:"thumbnails"`
}

type Thumbnails struct{
  Default Default `json:"default"`
}

type Default struct{
  Url string `json:"url"`
}
