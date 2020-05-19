package main

import (
	"skeleton/app"
)

func main() {
	service := app.NewPeopleServerService()
	service.Start()
}
