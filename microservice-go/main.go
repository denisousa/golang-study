package main

import (
	"productMs/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run() //running
}
