# sbglog
package for sbglog server



## UseConsole(out bool)
print messages to console

## UseSyslog(addr string)
print messages to syslog. for example __sbglog.UseSyslog("localhost:514")__

## UseGorutine(use bool)
use gorutine for sending messages

## Wait()
wait for gorutines

## SetName(name string)
set log name (tag for syslog)

## SetAddr(addr string)
addr for sbglog server. for example __sbglog.SetAddr("192.168.0.1:1514")__

## send log messages
* Emergency(i interface{})
* Alert(i interface{})
* Critical(i interface{})
* Error(i interface{})
* Warning(i interface{})
* Note(i interface{})
* Info(i interface{})
* Debug(i interface{})



