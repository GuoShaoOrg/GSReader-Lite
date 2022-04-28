package component

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

var logger *glog.Logger

func Logger() *glog.Logger {
	
	if logger != nil {
		return logger
	}

	logger = glog.New()
	_ = logger.SetConfigWithMap(g.Map{
		"path":     "/var/log",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
	})

	return logger
}
