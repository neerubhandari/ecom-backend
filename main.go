package main

import (
	"github.com/neerubhandari/ecom-backend/bootstrap"
	"github.com/neerubhandari/ecom-backend/utils"
)

func main() {
	utils.LoadEnv()
	bootstrap.WebApp()

}
