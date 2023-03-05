package module

import (
	"context"
	"course-explorer-monorepo/apps/api/content/core/repository"

	domain "github.com/gatherloop/course-explorer-monorepo"
)

type contentUsecase struct{
  contentRepo repository.ContentRepository
}

type ContentUsecase interface{
  GetCoursesList(ctx context.Context) ([]domain.GetCoursesList, error)
}

func NewContentUsecase(contentRepo repository.ContentRepository) ContentUsecase{
  return &contentUsecase{contentRepo}
}

func (c *contentUsecase) GetCoursesList(ctx context.Context) ([]domain.GetCoursesList, error){
  return c.contentRepo.GetCoursesList(ctx)
}
