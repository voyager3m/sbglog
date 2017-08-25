// sbglog project sbglog.go
package sbglog

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	name_    string = os.Args[0]
	addr_    string = "127.0.0.1:1514"
	conn_    net.Conn
	errtypes = []string{"EMERG", "ALERT", "CRITCL", "ERROR", "WARNG", "NOTE", "INFO", "DEBUG"}
)

func init() {
	var sp = strings.Split(name_, "/")
	name_ = sp[len(sp)-1]
	var err error
	conn_, err = net.Dial("udp", addr_)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connect to sbg server %v", err)
	}
}

func SetName(name string) {
	name_ = name
}

func SetAddr(addr string) {
	addr_ = addr
	vn := strings.Split(addr, ":")
	if len(vn) == 1 {
		addr_ += ":1514"
	}
	conn_.Close()
	var err error
	conn_, err = net.Dial("udp", addr_)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connect to sbg server %v", err)
	}
}

func Err(i interface{}) {
	s := fmt.Sprint(i)
	log(3, s)

}

func log(eventtype int, str string) {
	var _, file, line, _ = runtime.Caller(1)
	var pid = os.Getgid()
	var sf = strings.Split(file, "/")
	file = sf[len(sf)-1]
	t := time.Now().UTC()
	strt := t.Format(time.ANSIC)

	var fmtstr = fmt.Sprintf("<%d>%.15s %s[%d]: [%s:%d]: %s", eventtype, strt, name_, pid, file, line, str)
	fmt.Fprintf(conn_, fmtstr)

}
