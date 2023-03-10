package repository

import (
	"context"

	domain "github.com/gatherloop/course-explorer-monorepo"
	"gorm.io/gorm"
)

type instructorRepository struct {
	db *gorm.DB
}

type InstructorRepository interface {
	GetInstructorsList(ctx context.Context) ([]domain.Instructor, error)
}

func NewInstructorRepository(db *gorm.DB) InstructorRepository {
	return &instructorRepository{db: db}
}

func (i *instructorRepository) GetInstructorsList(ctx context.Context) ([]domain.Instructor, error) {
	var InstructorList []domain.Instructor
	err := i.db.Find(&InstructorList).Error
	return InstructorList, err

}
