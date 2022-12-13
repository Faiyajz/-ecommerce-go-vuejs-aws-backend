package main

import "fmt"

type User struct {
	ID      string
	Name    string
	Email   string
	Blocked bool
}

func blockUser(userPtr *User) {
	userPtr.Blocked = true
}

func main() {
	userInfo := User{
		ID:      "868",
		Name:    "Faiyaj",
		Email:   "faiyajz00@gmail.com",
		Blocked: false,
	}

	fmt.Println(userInfo)
	
	blockUser(&userInfo)
	fmt.Println(userInfo)
}
