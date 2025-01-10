package main

import (
	"BlogServer/utlis/pwd"
	"fmt"
)

func main() {
	p, _ := pwd.GenerateFromPassword("123456")
	fmt.Println(p)

	fmt.Println(pwd.CompareHashAndPassword(p, "123456"))
}
