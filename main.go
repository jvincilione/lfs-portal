package main

import (
	"fmt"
	"lfs-portal/routes"
	"lfs-portal/utils"

	"github.com/sirupsen/logrus"
)

func init() {
	utils.NewConfig()
}

func main() {
	router := routes.InitializeRoutes()
	logrus.Fatal(router.Run(fmt.Sprintf("0.0.0.0:%d", utils.Config.PORT)))
}
