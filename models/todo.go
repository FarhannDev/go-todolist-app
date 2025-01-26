package models

type Todo struct {
	ID      string `json:"id"`
	Task    string `json:"task"`
	Checked bool   `json:"checked"`
}

var Todos []Todo
