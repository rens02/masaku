package main

import (
	"masaku/config"
	"masaku/routes"
)

func main() {

	config.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
