package domain

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Age      int
}

func (User) ScName() string { return "User" }
