package main

import (
	"fmt"
	"lineNotify/routes"
)

func main() {
	fmt.Println("Server Launch!")

	// 註冊路由
	routes.Register()
}
