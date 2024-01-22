//go:build linux

package sys

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	ProcTcp6    = "/proc/net/tcp6"
	ProcTcp     = "/proc/net/tcp"
	ListenState = "0A"
)

// Read the table of tcp connections & remove header
func readFile(tcpfile string) []string {
	content, err := os.ReadFile(tcpfile)
	if err != nil {
		log.Fatalln(err, content)
	}
	return strings.Split(string(content), "\n")[1:]
}

func hexToDec(h string) int64 {
	dec, err := strconv.ParseInt(h, 16, 32)
	if err != nil {
		log.Fatalln(err)
	}
	return dec
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func Netstat(portToKill int64) []Process {
	tcpStats := statTCP(portToKill, ProcTcp)
	tcp6Stats := statTCP(portToKill, ProcTcp6)
	return append(tcpStats, tcp6Stats...)
}

// To get pid of all network process running on system, you must run this script
// as superuser
func statTCP(portToKill int64, tcpfile string) []Process {
	content := readFile(tcpfile)
	var processes []Process

	for _, line := range content {
		if line == "" {
			continue
		}
		parts := deleteEmpty(strings.Split(strings.TrimSpace(line), " "))
		localAddress := parts[1]
		state := parts[3]
		if state != ListenState {
			continue
		}
		inode := parts[9]
		localPort := hexToDec(strings.Split(localAddress, ":")[1])
		if localPort != portToKill {
			continue
		}

		pid := getPIDFromInode(inode)
		exe := getProcessExe(pid)
		p := Process{Name: exe, Pid: pid, State: state, Port: localPort}
		processes = append(processes, p)
	}

	return processes
}

// To retrieve the pid, check every running process and look for one using
// the given inode
func getPIDFromInode(inode string) string {
	pid := "-"

	d, err := filepath.Glob("/proc/[0-9]*/fd/[0-9]*")
	if err != nil {
		log.Fatalln(err)
	}

	re := regexp.MustCompile(inode)
	for _, item := range d {
		path, _ := os.Readlink(item)
		out := re.FindString(path)
		if len(out) != 0 {
			pid = strings.Split(item, "/")[2]
		}
	}
	return pid
}

func getProcessExe(pid string) string {
	exe := fmt.Sprintf("/proc/%s/exe", pid)
	path, _ := os.Readlink(exe)
	return path
}

func KillPort(portToKill int64) {
	killed := false
	for _, conn := range Netstat(portToKill) {
		if err := conn.Kill(); err != nil {
			log.Println(err)
		} else {
			log.Printf("Killed %s (pid: %s) listening on port %d", conn.Name, conn.Pid, conn.Port)
			killed = true
		}
	}
	if !killed {
		log.Printf("No process found listening on port %d\n", portToKill)
	}
}

func init() {
	log.SetFlags(0)
}
