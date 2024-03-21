package main

import (
	controller "exemple/controller"
	routes "exemple/routes"
)

func main() {
	controller.InitTemplates()
	routes.SetRoutes()
}
