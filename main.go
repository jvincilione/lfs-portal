package main

import (
	"lfs-portal/routes"
	"lfs-portal/utils"

	"github.com/sirupsen/logrus"
)

func init() {
	utils.NewConfig()
}

func main() {
	router := routes.InitializeRoutes()
	err := router.Run()
	if err != nil {
		logrus.Fatal("Error initializing routes.")
		return
	}
}
