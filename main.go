package main

import (
	"fmt"

	"github.com/Bo0km4n/go-mod-dummy/apis"
)

func main() {
	u := apis.User{
		Name: "kawabe",
	}

	fmt.Println(u)
}
