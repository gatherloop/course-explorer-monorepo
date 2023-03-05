package repository

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	domain "github.com/gatherloop/course-explorer-monorepo"
	"github.com/joho/godotenv"
)

var baseURL = "https://www.googleapis.com/youtube/v3"
var key = os.Getenv("API_KEY")
var channelId = os.Getenv("CHANNEL_ID")

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

type contentRepository struct{}

type ContentRepository interface{
  GetCoursesList(ctx context.Context) ([]domain.GetCoursesList, error)
}

func NewContentRepository() ContentRepository{
  return &contentRepository{}
}

func (repo *contentRepository) GetCoursesList(ctx context.Context) ([]domain.GetCoursesList, error){
  godotenv.Load()

  var getResponse GetResponse
  var courses []domain.GetCoursesList

  var client = &http.Client{}
  request, err := http.NewRequest("GET", baseURL+"/playlists?key="+key+"&part=snippet&channelId="+channelId+"&maxResults=30", nil)

  if err != nil {
    return []domain.GetCoursesList{}, err
  }

  response, err := client.Do(request)
  if err != nil {
    return []domain.GetCoursesList{}, err
  }
  defer response.Body.Close()

  err = json.NewDecoder(response.Body).Decode(&getResponse)
  if err != nil {
    return []domain.GetCoursesList{}, err
  }

  for _, playlist := range getResponse.Items{
    courses = append(courses, domain.GetCoursesList{Title: playlist.Snippet.Title, Description: playlist.Snippet.Description, Thumbnail: playlist.Snippet.Thumbnails.Default.Url})
  }

  return courses, nil
}
