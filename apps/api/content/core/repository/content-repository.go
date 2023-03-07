package repository

import (
	"context"
	"encoding/json"
	"net/http"

	course "course-explorer-monorepo/apps/api/content/core/entity"

	domain "github.com/gatherloop/course-explorer-monorepo"
)

type contentRepository struct{}

type ContentRepository interface{
  GetCoursesList(ctx context.Context) ([]domain.GetCoursesList, error)
}

func NewContentRepository() ContentRepository{
  return &contentRepository{}
}

func (repo *contentRepository) GetCoursesList(ctx context.Context) ([]domain.GetCoursesList, error){
  var getResponse course.GetResponse
  var courses []domain.GetCoursesList

  var client = &http.Client{}
  request, err := http.NewRequest("GET", course.BaseURL+"/playlists?key="+course.Key+"&part=snippet&channelId="+course.ChannelId+"&maxResults=30", nil)

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
