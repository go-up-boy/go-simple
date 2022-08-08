package app

import (
	"go-simple/pkg/config"
	"runtime"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

func TimeNowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}

func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}
