package main

import (
	"github.com/FishwithCat/gz-backend/src/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run()
}
