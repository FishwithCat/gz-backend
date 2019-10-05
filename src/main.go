package main

import (
	"github.com/FishwithCat/gz-backend/src/models"
	"github.com/FishwithCat/gz-backend/src/routers"
)

func init() {
	models.SetUp()
}

func main() {
	r := routers.InitRouter()
	r.Run()
}
