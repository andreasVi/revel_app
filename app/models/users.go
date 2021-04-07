package models

//Users
type User struct {
	ID       int    `form:"id" json:"id"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Name     string `form:"name" json:"name"`
	Age      int    `form:"age" json:"age"`
	Address  string `form:"address" json:"address"`
}

type UserResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}
