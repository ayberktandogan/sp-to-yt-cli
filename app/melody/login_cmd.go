package app

import "fmt"

type loginCmd struct {
}

func (i *loginCmd) Run() error {
	// TODO
	fmt.Println("login")

	return nil
}
