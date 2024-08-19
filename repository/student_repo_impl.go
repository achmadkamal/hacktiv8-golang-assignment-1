package repository

import (
	"errors"
	"hacktiv8_golang_assignment_1/entity"
)

type StudentRepoImpl struct {
	students []entity.Student
}

func NewStudentRepo(students []entity.Student) StudentRepo {
	return &StudentRepoImpl{
		students: students,
	}
}

func (repo StudentRepoImpl) FindByName(name string) (*entity.Student, error) {
	for _, student := range repo.students {
		if student.Name == name {
			return &student, nil
		}
	}
	return nil, errors.New("data not found!")
}
