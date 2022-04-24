package server

import (
	"gs-reader-lite/server/routers"
)

func Run() {
	routers.InitRouter()
}