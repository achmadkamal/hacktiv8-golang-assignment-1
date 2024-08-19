package main

import (
	"fmt"
	"hacktiv8_golang_assignment_1/data"
	"hacktiv8_golang_assignment_1/repository"
	"os"
)

func main() {
	name := os.Args[1]

	dataStudent := data.Students
	studentRepo := repository.NewStudentRepo(dataStudent)

	student, err := studentRepo.FindByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Id: %d\nName: %s\nAddress: %s\nJob: %s\nReason: %s\n",
		student.Id,
		student.Name,
		student.Address,
		student.Job,
		student.Reason,
	)
}
