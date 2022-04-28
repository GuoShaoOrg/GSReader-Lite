package server

import (
	"gs-reader-lite/server/component"
	"gs-reader-lite/server/routers"
)

func Run() {
	component.InitDatabase()
	routers.InitRouter()
}