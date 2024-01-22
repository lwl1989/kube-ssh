//go:build windows

package sys

import (
	"fmt"
	"strconv"
	"syscall"
	"unsafe"
)

func KillPort(p int64) {
	res := Netstat(p)
	for _, v := range res {
		_ = v.Kill()
	}
}

func Netstat(port int64) (res []Process) {
	lazydll := syscall.NewLazyDLL("Iphlpapi.dll")
	proc := lazydll.NewProc("GetTcpTable2")

	var mibtable2 MIB_TCPTABLE2
	size := unsafe.Sizeof(mibtable2)

	//第一次执行是获取缓存区大小,然后根据返回的size申请对应长度的内存
	r, _, err := proc.Call(uintptr(unsafe.Pointer(&mibtable2)), uintptr(unsafe.Pointer(&size)), 1)
	if err != nil && r != 0 {
		if r == ERROR_INSUFFICIENT_BUFFER {
			buf := make([]byte, size)
			r, _, err = proc.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&size)), 1)
			if r != 0 {
				fmt.Printf("Get tcp table error:%s\n", err.Error())
				return
			}
			var index = int(unsafe.Sizeof(mibtable2.dwNumEntries))
			var step = int(unsafe.Sizeof(mibtable2.table))
			dwNumEntries := *(*uint32)(unsafe.Pointer(&buf[0]))
			for i := 0; i < int(dwNumEntries); i++ {
				mibs := *(*MIB_TCPROW2)(unsafe.Pointer(&buf[index]))
				index += step
				if mibs.dwState != 2 {
					continue
				}
				portStr := mibs.dwLocalPort.String()
				listenP, _ := strconv.Atoi(portStr)
				if int64(listenP) == port {
					res = append(res, Process{
						Name:  mibs.String(),
						Pid:   strconv.FormatInt(int64(mibs.dwOwningPid), 10),
						State: mibs.dwState.String(),
						Port:  int64(mibs.dwLocalPort),
					})
				}
			}
		}
	}
	return nil
}

func NetStatWithPid(pid int) *Process {
	return nil
}

const ERROR_INSUFFICIENT_BUFFER = 122

type inet_ntoa uint32

// 地址转化
func (i inet_ntoa) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i&255, i>>8&255, i>>16&255, i>>24&255)
}

type ntohs uint32

// 端口转化
func (i ntohs) String() string {
	return fmt.Sprint(syscall.Ntohs(uint16(i)))
}

type TCP_CONNECTION_OFFLOAD_STATE uint32

// 状态枚举
var _MIB_TCP_STATE = map[uint32]string{
	1:  "CLOSED",
	2:  "LISTEN",
	3:  "SYN_SENT",
	4:  "SYN_RCVD",
	5:  "ESTABLISHED",
	6:  "FIN_WAIT1",
	7:  "FIN_WAIT2",
	8:  "CLOSE_WAIT",
	9:  "CLOSING",
	10: "LAST_ACK",
	11: "TIME_WAIT",
	12: "DELETE_TCB",
}

type MIB_TCP_STATE uint32

func (m MIB_TCP_STATE) String() string {
	return _MIB_TCP_STATE[uint32(m)]
}

type MIB_TCPROW2 struct {
	dwState        MIB_TCP_STATE
	dwLocalAddr    inet_ntoa
	dwLocalPort    ntohs
	dwRemoteAddr   inet_ntoa
	dwRemotePort   ntohs
	dwOwningPid    uint32
	dwOffloadState TCP_CONNECTION_OFFLOAD_STATE
}

func (M MIB_TCPROW2) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%d", M.dwLocalAddr, M.dwLocalPort, M.dwRemoteAddr, M.dwRemotePort, M.dwState, M.dwOwningPid)
}

type MIB_TCPTABLE2 struct {
	dwNumEntries uint32
	table        [1]MIB_TCPROW2
}

/*
const (
    //MIB_TCPTABLE
    TCP_TABLE_BASIC_LISTENER uintptr = iota + 1
    TCP_TABLE_BASIC_CONNECTIONS
    TCP_TABLE_BASIC_ALL

    //MIB_TCPTABLE_OWNER_PID
    TCP_TABLE_OWNER_PID_LISTENER
    TCP_TABLE_OWNER_PID_CONNECTIONS
    TCP_TABLE_OWNER_PID_ALL

    //MIB_TCPTABLE_OWNER_MODULE
    TCP_TABLE_OWNER_MODULE_LISTENER
    TCP_TABLE_OWNER_MODULE_CONNECTIONS
    TCP_TABLE_OWNER_MODULE_ALL
)

const (
    AF_INET  = 2
    AF_INET6 = 23
)

type MIB_TCPROW_OWNER_PID struct {
    dwState      MIB_TCP_STATE
    dwLocalAddr  inet_ntoa
    dwLocalPort  ntohs
    dwRemoteAddr inet_ntoa
    dwRemotePort ntohs
    dwOwningPid  uint32
}

func (M MIB_TCPROW_OWNER_PID) String() string {
    return fmt.Sprintf("%s %s %s %s %s %d", M.dwLocalAddr, M.dwLocalPort, M.dwRemoteAddr, M.dwRemotePort, M.dwState, M.dwOwningPid)
}

func MIB_TCPROW_OWNER_PID_TOSTRING(p unsafe.Pointer) fmt.Stringer {
    return *(*MIB_TCPROW_OWNER_PID)(p)
}

type MIB_TCPTABLE_OWNER_PID struct {
    dwNumEntries uint32
    table        [1]MIB_TCPROW_OWNER_PID
}
*/
