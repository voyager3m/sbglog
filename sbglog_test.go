package sbglog

import (
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	//UseConsole(true)
	Infof("hello world %d", 12)
	time.Sleep(time.Second)
}
