package model

type Task struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}