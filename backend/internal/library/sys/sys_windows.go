//go:build windows
// +build windows

package sys

func getPidFile() string {
	return ""
}

func WritePidFile() {

}

func RmPidFile() {

}
