# sbglog
package for sbglog server



## UseConsole(out bool)
print messages to console

## UseSyslog(addr string)
print messages to syslog. for example __sbglog.UseSyslog("localhost:514")__


## SetName(name string)
set log name (tag for syslog)

## SetAddr(addr string)
addr for sbglog server. for example __sbglog.SetAddr("192.168.0.1:1514")__

## SetLogLevel(loglevel int)
set log level filter from 0 (nothing out) to 7 (debug)

## Check(err error)
check and return true if error == nil or log error

## send log messages
* Emergency(i interface{})
* EmergencyWait(i interface{})
* Alert(i interface{})
* AlertWait(i interface{})
* Critical(i interface{})
* CriticalWait(i interface{})
* Error(i interface{})
* ErrorWait(i interface{})
* Warning(i interface{})
* WarningWait(i interface{})
* Note(i interface{})
* NoteWait(i interface{})
* Info(i interface{})
* InfoWait(i interface{})
* Debug(i interface{})
* DebugWait(i interface{})



