//go:build !windows && !linux

package sys

func Netstat(portToKill int64) []Process {
	return nil
}

func KillPort(portToKill int64) {

}
