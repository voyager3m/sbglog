// sbglog project sbglog.go
package sbglog

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	name_        string = os.Args[0]
	addr_        string = "127.0.0.1:1514"
	conn_        net.Conn
	errtypes          = []string{"EMERG", "ALERT", "CRITCL", "ERROR", "WARNG", "NOTE", "INFO", "DEBUG"}
	usegorutine_ bool = false
	consoleout_  bool = false
	syslogout_   bool = false
	wg_          sync.WaitGroup
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

func SetConsoleOut(out bool) {
	consoleout_ = out
}

func SetSyslogOut(out bool) {
	syslogout_ = out
}

func UseGorutine(use bool) {
	usegorutine_ = use
}

func SetName(name string) {
	name_ = name
}

func Wait() {
	wg_.Wait()
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

func Emergency(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go log(0, file, line, s)
	} else {
		log(0, file, line, s)
	}

}

func Alert(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go log(1, file, line, s)
	} else {
		log(1, file, line, s)
	}

}

func Critical(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go log(2, file, line, s)
	} else {
		log(2, file, line, s)
	}
}

func Error(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go log(3, file, line, s)
	} else {
		log(3, file, line, s)
	}
}

func Warning(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go log(4, file, line, s)
	} else {
		log(4, file, line, s)
	}
}

func Note(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(2)
	if usegorutine_ {
		wg_.Add(1)
		go log(5, file, line, s)
	} else {
		log(5, file, line, s)
	}
}

func Info(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(2)
	if usegorutine_ {
		wg_.Add(1)
		go log(6, file, line, s)
	} else {
		log(6, file, line, s)
	}
}

func Debug(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(2)
	if usegorutine_ {
		wg_.Add(1)
		go log(7, file, line, s)
	} else {
		log(7, file, line, s)
	}
}

func log(eventtype int, file string, line int, str string) {
	var pid = syscall.Getpid() //os.Getgid()
	var sf = strings.Split(file, "/")
	file = sf[len(sf)-1]
	t := time.Now().UTC()
	strt := t.Format("Jan 2 15:04:05 2006")

	var fmtstr = fmt.Sprintf("<%d>%.15s %s[%d]: %s:%d %s: %s", eventtype, strt, name_, pid, file, line, errtypes[eventtype], str)
	fmt.Fprintf(conn_, fmtstr)
	if consoleout_ {
		fmt.Fprintf(os.Stderr, "%s\n", fmtstr)
	}
	if syslogout_ {

	}
	if usegorutine_ {
		wg_.Done()
	}
}
