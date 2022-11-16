package main

import (
	"mall/conf"
	"mall/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	r.Run(":8081")
}
