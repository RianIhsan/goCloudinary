package models

type User struct {
	Name  string `json:"name" form:"name"`
	Image string `json:"image" form:"image"`
}
