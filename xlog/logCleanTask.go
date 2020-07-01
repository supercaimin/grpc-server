/* logCleanTask.go
作者：	汪军
日期：2019-9-4

日志清理操作
*/
package xlog

import (
	"compress/gzip"

	"fmt"
	"os"
	"time"

	"hstcmler/errcode"
)

type tLogFileStats struct {
	fileName        string
	cumulativeBytes int64 //累计的字节数
}

const MAX_LOG_FILE int = 365

var (
	gFileStats [MAX_LOG_FILE]tLogFileStats
	logserv    *TLogService
	fileNum    int
)

/*初始化Log清理任务*/
func InitLogCleanTask(parent *TLogService) {
	logserv = parent
	go LogCleanTask() //日志清理的任务
}

/*整理日志文件，仅保留最近30天并不超过总容量的日志文件，并且对每天的日志文件进行压缩处理
对于一天以及以前的文件压缩为.zip文件，并删除原始Log文件*/
func LogCleanTask() {

	var (
		f1, f2    string
		err       error
		fi        os.FileInfo
		toBeClean bool = false
	)

	time.Sleep(2 * time.Second)
	fmt.Println("[LogClean Task]Begin Log Clean Task")
	time.AfterFunc(time.Hour*1, LogCleanTask)
	fileNum = 0

	cur := time.Now()
	if err != nil {
		fmt.Printf("[LogClean Task]List Log file Dir Failed, err:%s\n", err)
		return
	}

	for i := 1; i < MAX_LOG_FILE; i++ {
		tm := cur.Add(-1 * time.Hour * 24 * time.Duration(i))
		f1 = getLogFileName(&tm)
		f2 = logserv.GetLogFilePath() + f1
		f1 = f2 + ".zip"
		if toBeClean {
			os.Remove(f1)
			os.Remove(f2)
			continue
		}

		fi, err = os.Stat(f2)
		if err == nil {
			if errcode.RESULT_SUCCESS == compressFile(f2, f1) {
				os.Remove(f2)
			} else {
				continue
			}

		}

		fi, err = os.Stat(f1)
		if err == nil {
			if fileNum <= 0 {
				gFileStats[0].cumulativeBytes = fi.Size()
			} else {
				gFileStats[fileNum].cumulativeBytes = gFileStats[fileNum-1].cumulativeBytes + fi.Size()
			}
			if (gFileStats[fileNum].cumulativeBytes > logserv.maxStorageSize) || (i > logserv.maxStoragePeriod) {
				toBeClean = true
			}

			fileNum++
		}
	}

}

/*压缩一个文件，用gzip算法进行压缩*/
func compressFile(fileIn string, fileOut string) errcode.RESULT {

	var (
		fp1, fp2 *os.File
		err      error
		buf      [8192]byte
		n1       int
	)

	fp1, err = os.Open(fileIn)
	if err != nil {
		return errcode.RESULT_ERROR_FAILED_OPEN_FILE
	}

	fp2, err = os.Create(fileOut)
	if err != nil {
		fp1.Close()
		fmt.Printf("[LogCleanTask]Create Compressed Target File Failed, name=%s, err=%s\n", fileOut, err)
		return errcode.RESULT_ERROR_FAILED_OPEN_FILE

	}

	gw := gzip.NewWriter(fp2)

	for n1, err = fp1.Read(buf[0:]); n1 > 0; n1, err = fp1.Read(buf[0:]) {
		gw.Write(buf[0:n1])
	}
	gw.Flush()
	gw.Close()
	fp1.Close()
	fp2.Close()

	return errcode.RESULT_SUCCESS

}
