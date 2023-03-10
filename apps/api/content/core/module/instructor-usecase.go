package module

import (
	"context"
	"course-explorer-monorepo/apps/api/content/core/repository"

	domain "github.com/gatherloop/course-explorer-monorepo"
)

type instructorUsecase struct {
	instructorRepo repository.InstructorRepository
}

type InstructorUsecase interface {
	GetInstructorsList(ctx context.Context) ([]domain.Instructor, error)
}

func NewInstructorUsecase(instructorRepo repository.InstructorRepository) InstructorUsecase {
	return &instructorUsecase{instructorRepo: instructorRepo}
}
func (i *instructorUsecase) GetInstructorsList(ctx context.Context) ([]domain.Instructor, error) {
	return i.instructorRepo.GetInstructorsList(ctx)
}
