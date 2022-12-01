package models

type Student struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
}

var Students = []Student{
	Student{ID: 1, FullName: "Pedro", Age: 03},
	Student{ID: 2, FullName: "Amanda", Age: 24},
}
