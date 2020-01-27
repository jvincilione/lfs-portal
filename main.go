package main

import (
	"lfs-portal/routes"

	"github.com/sirupsen/logrus"
)

func main() {
	router := routes.InitializeRoutes()
	err := router.Run()
	if err != nil {
		logrus.Fatal("Error initializing routes.")
		return
	}
}
