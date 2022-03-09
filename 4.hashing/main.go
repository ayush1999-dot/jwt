package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "ayush"
	bs, err := hashingpass(pass)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	err = check("shetty", bs)
	if err != nil {
		fmt.Println("ur not logged in")
	} else {
		fmt.Println("logged in")
	}
}
func hashingpass(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return bs, nil
}
func check(password string, bs []byte) error {
	err := bcrypt.CompareHashAndPassword(bs, []byte(password))

	return err

}
