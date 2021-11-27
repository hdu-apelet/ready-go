package internal

import (
	_ "readygo/internal/database"
	"readygo/internal/route"
	_ "readygo/internal/route/api"
)

func Run() {
	route.Run(":8080")
}
