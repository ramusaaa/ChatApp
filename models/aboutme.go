package models

type Aboutme struct {
	Id        uint   `json:"id" gorm:"primaryKey" gorm:"unique"`
	Image     string `json:"image"`
	Content   string `json:"content"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
