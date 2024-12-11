package main

import (
	"Blog/core"
	"fmt"
)

func main() {
	core.InitIPDB()
	fmt.Println(core.GetIPAddr("175.0.201.207"))
	fmt.Println(core.GetIPAddr("127.0.0.1"))
	fmt.Println(core.GetIPAddr("10.0.201.207"))
	fmt.Println(core.GetIPAddr("8.4.5.6"))
	fmt.Println(core.GetIPAddr("223.104.194.176"))

}
