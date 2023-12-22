package app

import (
	"fmt"
)

type loginCmd struct {
}

func (i *loginCmd) Run() error {
	Open_Url("https://example.com")
	fmt.Println("login")

	return nil
}
