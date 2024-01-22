//go:build linux

package sys

import (
	"backend-go/api/internal/config"
	"os"
	"strconv"
)

func getPidFile() string {
	name := config.GlobalConfig.Name
	pidFile := "/run/" + name + ".pid"
	if os.Getenv("GO_ENV") != "" {
		pidFile = "/run/" + name + "-" + os.Getenv("GO_ENV") + ".pid"
	}
	return pidFile
}

func WritePidFile() {
	pidFile := getPidFile()
	pidStr := strconv.FormatInt(int64(os.Getpid()), 10)
	_ = os.WriteFile(pidFile, []byte(pidStr), 0777)
}

func RmPidFile() {
	pidFile := getPidFile()
	_ = os.Remove(pidFile)
}
