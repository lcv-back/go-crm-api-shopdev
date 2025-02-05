package main

import (
	"github.com/lcv-back/go-crm-api-shopdev/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8002")
}
