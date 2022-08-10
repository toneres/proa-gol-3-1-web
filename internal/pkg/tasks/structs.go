package tasks

import "github.com/toneres/proa-gol-3-1-web/internal/pkg/employees"

type Task struct {
	ID        int64
	Name      string `validate:"required,min=3,max=160"`
	Deadline  string `validate:"required,datetime=2006-01-02"`
	Status    string `validate:"required,oneof=open progress done"`
	Assignees []employees.Employee
}

type TaskAttributes struct {
	ID       int64
	Name     string `validate:"required,min=3,max=160"`
	Deadline string `validate:"required,datetime=2006-01-02"`
	Status   string `validate:"required,oneof=open progress done"`
}
