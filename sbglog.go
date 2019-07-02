// sbglog project sbglog.go
package sbglog

import (
	"fmt"
	"log"
	"log/syslog"
	"net"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"
)

var (
	name_       string = os.Args[0]
	addr_       string //= "127.0.0.1:1514"
	conn_       net.Conn
	connected_  bool = false
	errtypes         = []string{"EMERG", "ALERT", "CRITCL", "ERROR", "WARNG", "NOTE", "INFO", "DEBUG"}
	consoleout_ bool = true
	syslogout_  bool = false
	syslog_     *syslog.Writer
	loglevel_   int = 7
	pid_        int
)

func init() {
	var sp = strings.Split(name_, "/")
	name_ = sp[len(sp)-1]
	// var err error
	// conn_, err = net.Dial("udp", addr_)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error connect to sbg server %v", err)

	// }
	pid_ = syscall.Getpid()
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

func SetName(name string) {
	name_ = name
}

// SetAddr set url to sbglog server
func SetAddr(addr string) {
	addr_ = addr
	if connected_ {
		conn_.Close()
	}
	connected_ = false
	if len(addr_) > 0 {
		vn := strings.Split(addr, ":")
		if len(vn) == 1 {
			addr_ += ":1514"
		}
		var err error
		conn_, err = net.Dial("udp", addr_)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error connect to sbg server %v", err)
		} else {
			connected_ = true
		}
	}
}

/*set max loglevel , from 0 (nothing) to 7(debug) */
func SetLogLevel(level int) {
	loglevel_ = level
}

func Emergency(i interface{}) {
	if loglevel_ >= 0 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(0, file, line, s)
		go vlog(0, file, line, s)
	}
}

func EmergencyWait(i interface{}) {
	if loglevel_ >= 0 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(0, file, line, s)
		vlog(0, file, line, s)
	}
}

func Alert(i interface{}) {
	if loglevel_ >= 1 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(1, file, line, s)
		go vlog(1, file, line, s)
	}
}

func AlertWait(i interface{}) {
	if loglevel_ >= 1 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(1, file, line, s)
		vlog(1, file, line, s)
	}
}

func Critical(i interface{}) {
	if loglevel_ >= 2 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(2, file, line, s)
		go vlog(2, file, line, s)
	}
}

func CriticalWait(i interface{}) {
	if loglevel_ >= 2 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(2, file, line, s)
		vlog(2, file, line, s)
	}
}

func Error(i interface{}) {
	if loglevel_ >= 3 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(3, file, line, s)
		go vlog(3, file, line, s)
	}
}

func ErrorWait(i interface{}) {
	if loglevel_ >= 3 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(3, file, line, s)
		vlog(3, file, line, s)
	}
}

func Warning(i interface{}) {
	if loglevel_ >= 4 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(4, file, line, s)
		go vlog(4, file, line, s)
	}
}

func WarningWait(i interface{}) {
	if loglevel_ >= 4 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(4, file, line, s)
		vlog(4, file, line, s)
	}
}

func Note(i interface{}) {
	if loglevel_ >= 5 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(5, file, line, s)
		go vlog(5, file, line, s)
	}
}

func Notef(format string, i ...interface{}) {
	if loglevel_ >= 5 {
		s := fmt.Sprintf(format, i...)
		var _, file, line, _ = runtime.Caller(1)
		conlog(5, file, line, s)
		go vlog(5, file, line, s)
	}
}

func NoteWait(i interface{}) {
	if loglevel_ >= 5 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(5, file, line, s)
		vlog(5, file, line, s)
	}
}

func Info(i interface{}) {
	if loglevel_ >= 6 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(6, file, line, s)
		go vlog(6, file, line, s)
	}
}

func Infof(format string, i ...interface{}) {
	if loglevel_ >= 6 {
		s := fmt.Sprintf(format, i...)
		var _, file, line, _ = runtime.Caller(1)
		conlog(6, file, line, s)
		go vlog(6, file, line, s)
	}
}

func InfoWait(i interface{}) {
	if loglevel_ >= 6 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(6, file, line, s)
		vlog(6, file, line, s)
	}
}

func Debug(i interface{}) {
	if loglevel_ >= 7 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(7, file, line, s)
		go vlog(7, file, line, s)
	}
}

func Debugf(format string, i ...interface{}) {
	if loglevel_ >= 7 {
		s := fmt.Sprintf(format, i...)
		var _, file, line, _ = runtime.Caller(1)
		conlog(7, file, line, s)
		go vlog(7, file, line, s)
	}
}

func DebugWait(i interface{}) {
	if loglevel_ >= 0 {
		s := fmt.Sprint(i)
		var _, file, line, _ = runtime.Caller(1)
		conlog(7, file, line, s)
		vlog(7, file, line, s)
	}
}

func DebugfWait(format string, i ...interface{}) {
	if loglevel_ >= 0 {
		s := fmt.Sprintf(format, i...)
		var _, file, line, _ = runtime.Caller(1)
		conlog(7, file, line, s)
		vlog(7, file, line, s)
	}
}

func Check(err error) bool {
	if err != nil {
		if loglevel_ >= 3 {
			var _, file, line, _ = runtime.Caller(1)
			s := fmt.Sprintf("CHECK: %+v", err)
			conlog(3, file, line, s)
			vlog(3, file, line, s)
		}
		return false
	}
	return true
}

func vlog(eventtype int, file string, line int, str string) {
	//conlog(eventtype, file, line, str)
	if !connected_ && !syslogout_ {
		return
	}
	var sf = strings.Split(file, "/")
	file = sf[len(sf)-1]
	t := time.Now().UTC()
	strt := t.Format("Jan 2 15:04:05 2006")
	var fmtstr = fmt.Sprintf("<%d>%.15s %s[%d]: %s:%d %s: %s", eventtype, strt, name_, pid_, file, line, errtypes[eventtype], str)
	if connected_ {
		fmt.Fprintf(conn_, fmtstr)
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
}

func conlog(eventtype int, file string, line int, str string) {
	if consoleout_ {
		var sf = strings.Split(file, "/")
		log.Println(fmt.Sprintf("%s (%s:%d): %s", errtypes[eventtype], sf[len(sf)-1], line, str))
	}
}
