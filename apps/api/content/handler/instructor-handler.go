package handler

import (
	"course-explorer-monorepo/apps/api/content/core/module"
	"course-explorer-monorepo/libs/api/utils/response"
	"net/http"

	domain "github.com/gatherloop/course-explorer-monorepo"
)

type instructorHandler struct {
	instructorUsecase module.InstructorUsecase
}

type InstructorHandler interface {
	GetInstructorsList(w http.ResponseWriter, r *http.Request)
}

func NewInstructorHandler(instructorUsecase module.InstructorUsecase) InstructorHandler {
	return &instructorHandler{instructorUsecase: instructorUsecase}
}

func (i *instructorHandler) GetInstructorsList(w http.ResponseWriter, r *http.Request) {
	instructors, err := i.instructorUsecase.GetInstructorsList(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, nil, err.Error())
	}
	var ResponseGetInstructorsList domain.GetInstructorsListResponse
	ResponseGetInstructorsList.SetMessage("Succes Get Data")
	ResponseGetInstructorsList.SetData(instructors)
	response.Success(w, http.StatusOK, ResponseGetInstructorsList)
}
