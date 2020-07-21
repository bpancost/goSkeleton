package main

import (
	"goSkeleton/app"
)

func main() {
	service := app.NewPeopleServerService()
	service.Start()
}
