package main

import (
	"DineEasy/model"
	"DineEasy/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
