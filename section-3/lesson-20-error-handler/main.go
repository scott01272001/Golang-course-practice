package main

import (
	"fmt"
)

type errUserNameExist struct {
	UserName string
}

func (e errUserNameExist) Error() string {
	return fmt.Sprintf("username %s alerady exist.", e.UserName)
}

func isErrUserNameExist(err error) bool {
	// v, ok = i.(T)
	// interface type assert.
	// cast an interface into a concrete type, i is interace, T is a concrete type.
	// if the T's underlying type is not i then return false to ok.
	_, ok := err.(errUserNameExist)
	return ok
}

func checkUserNameExist(username string) (bool, error) {
	if username == "appleboy" {
		return true, errUserNameExist{UserName: username}
	}
	return false, nil
}

func main() {
	if _, err := checkUserNameExist("appleboy"); err != nil {
		if ok := isErrUserNameExist(err); ok {
			fmt.Println("appleboy error exist.")
		}
	}
}
