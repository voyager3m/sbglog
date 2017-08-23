// sbglog project sbglog.go
package sbglog

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
)

var (
	name_ string = os.Args[0]
	addr_ string = "127.0.0.1:1514"
)

func init() {
	var sp = strings.Split(name_, "/")
	name_ = sp[len(sp)-1]
}

func SetName(name string) {
	name_ = name
}

func SetAddr(addr string) {
	addr_ = addr
}

func Log(str string) {
	conn, err := net.Dial("udp", addr_)
	if err != nil {
		fmt.Printf("Error connect to sbg server %v", err)
	}
	defer conn.Close()
	var _, file, line, _ = runtime.Caller(1)
	var pid = os.Getgid()
	var sf = strings.Split(file, "/")
	file = sf[len(sf)-1]
	var fmtstr = fmt.Sprintf("%s[%d]: [%s:%d]: %s", name_, pid, file, line, str)
	fmt.Fprintf(conn, fmtstr)

}
