package repository

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserRepository interface {
	FindByName(name string) *User
	Create(name string, age int) *User
	DeleteByName(name string) *User
}
