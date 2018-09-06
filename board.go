package main

type Board struct {
	UniqueId string `json:"unique_id"`
	Password string `json:"password"`
	Tasks    []Task `json:"tasks"`
}
