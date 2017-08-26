// sbglog project sbglog.go
package sbglog

import (
	"fmt"
	"log/syslog"
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
	syslog_      *syslog.Writer
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

func UseConsole(out bool) {
	consoleout_ = out
}

func UseSyslog(addr string) {
	if syslogout_ {
		syslog_.Close()
	}
	if len(addr) > 0 {
		var err error
		syslog_, err = syslog.Dial("udp", addr, syslog.LOG_DEBUG|syslog.LOG_USER, name_)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error dial syslog: %v", err)
		}
		syslogout_ = true
	} else {
		syslogout_ = false
	}

}

func UseGorutine(use bool) {
	usegorutine_ = use
}

func Wait() {
	wg_.Wait()
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

func Emergency(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go vlog(0, file, line, s)
	} else {
		vlog(0, file, line, s)
	}

}

func Alert(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go vlog(1, file, line, s)
	} else {
		vlog(1, file, line, s)
	}

}

func Critical(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go vlog(2, file, line, s)
	} else {
		vlog(2, file, line, s)
	}
}

func Error(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go vlog(3, file, line, s)
	} else {
		vlog(3, file, line, s)
	}
}

func Warning(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go vlog(4, file, line, s)
	} else {
		vlog(4, file, line, s)
	}
}

func Note(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go vlog(5, file, line, s)
	} else {
		vlog(5, file, line, s)
	}
}

func Info(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go vlog(6, file, line, s)
	} else {
		vlog(6, file, line, s)
	}
}

func Debug(i interface{}) {
	s := fmt.Sprint(i)
	var _, file, line, _ = runtime.Caller(1)
	if usegorutine_ {
		wg_.Add(1)
		go vlog(7, file, line, s)
	} else {
		vlog(7, file, line, s)
	}
}

func vlog(eventtype int, file string, line int, str string) {
	var pid = syscall.Getpid()
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
		switch eventtype {
		case 0:
			syslog_.Emerg(fmtstr)
		case 1:
			syslog_.Alert(fmtstr)
		case 2:
			syslog_.Crit(fmtstr)
		case 3:
			syslog_.Err(fmtstr)
		case 4:
			syslog_.Warning(fmtstr)
		case 5:
			syslog_.Notice(fmtstr)
		case 6:
			syslog_.Info(fmtstr)
		case 7:
			syslog_.Debug(fmtstr)
		}
	}
	if usegorutine_ {
		wg_.Done()
	}
}
