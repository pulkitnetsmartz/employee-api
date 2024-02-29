package models

type Employee_detail struct {
	EmployeeId int    `json:"id"`
	Name       string `json:"firstname"`
	Role       int    `json:"role"`
	ProjectID  int    `json:"project_id"`
}

type Project_detail struct {
	ProjectID   int    `json:"project_id"`
	ProjectName string `json:"project_name"`
}
