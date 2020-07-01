// log_test.go
package xlog

import (
	"fmt"
	"testing"
	"time"

	"hstcmler/errcode"
)

var log *xLogger

func TestUtils(t *testing.T) {
	var i int
	ip := "127.0.0.1"
	ret := InitLogService(ip, 10000, "d:\\develop\\testfile", "anpc")
	if ret != errcode.RESULT_SUCCESS {
		t.Fatalf("Failed init Log Service, err=%s", ret.String())
	}
	SetModulePersitentLogLevel("test", LOG_LEVEL_INFO)

	PrintLogServiceStats()

	time.Sleep(time.Millisecond * 200)
	log = NewXLogger("test", "0.0.0.0", 0, ip, 10000)
	if log == nil {
		t.Fatalf("Error while Init log client")
	}

	for i = 0; i < 2048; i++ {
		log.Errorf("testmain", errcode.RESULT_ERROR_CPU_OVERLOAD, "CPU overload")

		log.Infof("testmain", errcode.RESULT_SUCCESS, "Info Test, Seq=%d", i)
		log.Debugf("testmain", errcode.RESULT_SUCCESS, "Debug test,i=%d", i)
		time.Sleep(1 * time.Millisecond)
	}

	log.PrintStats()
	PrintLogServiceStats()
	time.Sleep(30 * time.Second)
	fmt.Println("--------------Ultimate Result------------")
	log.PrintStats()
	PrintLogServiceStats()

}
