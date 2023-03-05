package handler

import (
	"course-explorer-monorepo/apps/api/content/core/module"
	"course-explorer-monorepo/libs/api/utils/response"
	"net/http"

	domain "github.com/gatherloop/course-explorer-monorepo"
)

type contentHandler struct{
  contentUsecase module.ContentUsecase
}

type ContentHandler interface{
  GetContentList(w http.ResponseWriter, r *http.Request)
}

func NewContentHandler(contentUsecase module.ContentUsecase) ContentHandler{
  return &contentHandler{contentUsecase}
}

func (c *contentHandler) GetContentList(w http.ResponseWriter, r *http.Request){
  courses, err := c.contentUsecase.GetCoursesList(r.Context())
  if err != nil{
    response.Error(w, http.StatusInternalServerError, nil, err.Error())
  }

  var getCoursesListResponse domain.GetCoursesListResponse
  getCoursesListResponse.SetMessage("success get courses list data")
  getCoursesListResponse.SetData(courses)
  response.Success(w, http.StatusOK, getCoursesListResponse)
}
