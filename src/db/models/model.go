package models

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey ;autoIncrement"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Desc   string `json:"desc"`
}

type User struct {
	Id       int    `json:"id" gorm:"primaryKey ;autoIncrement" `
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"biography"`
}
