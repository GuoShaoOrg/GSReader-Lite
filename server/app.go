package server

import (
	"gs-reader-lite/server/api/jobs"
	"gs-reader-lite/server/component"
	"gs-reader-lite/server/routers"
)

func Run() {
	component.InitDatabase()
	jobs.RegisterJob()
	routers.InitRouter()
}