package models

import "github.com/k0kubun/pp"

type User struct {
	Address string
}

func CreateUser(email string) User {
	return User{Address: email}
}

func handleNewAddressRequest(s string) (string, string) {
	pp.Println(s)
	return "", ""
}

func addressFromMail(s string) string {
	return s
}

func visitLinkFromEmailTo(s string) error {
	pp.Println(s)
	return nil
}

func addressFromDB(user User) string {
	return user.Address
}
