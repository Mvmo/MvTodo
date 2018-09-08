package main

type Board struct {
	Password string `json:"password"`
	Tasks    []Task `json:"tasks"`
}
