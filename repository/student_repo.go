package repository

import "hacktiv8_golang_assignment_1/entity"

type StudentRepo interface {
	FindByName(name string) (*entity.Student, error)
}
