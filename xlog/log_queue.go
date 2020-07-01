/* log_queue.go
作者：	汪军
日期：2019-9-4

日志队列操作
*/
package xlog

import (
	"fmt"
	"sync"

	"hstcmler/errcode"
)

type LogQueue struct {
	logQueue   []*InnerLogMsg
	cap        int
	queue_len  int
	head, rear int
	mutex      sync.Mutex
}

/*创建一个Queue*/
func NewLogQueue(capacity int) *LogQueue {

	q := new(LogQueue)
	q.logQueue = make([]*InnerLogMsg, capacity)
	q.cap = capacity
	q.Clear()

	return q
}

/*日志入队，压入队尾*/
func (self *LogQueue) PushBack(msg *InnerLogMsg) errcode.RESULT {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	if self.queue_len >= self.cap {
		return errcode.RESULT_ERROR_OUTOF_MEMORY
	}
	self.rear = (self.rear + 1) % self.cap

	self.logQueue[self.rear] = msg

	self.queue_len++
	return errcode.RESULT_SUCCESS
}

/*日志出队，始终从队列头部出队*/
func (self *LogQueue) Pop() *InnerLogMsg {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	if self.queue_len <= 0 {
		return nil
	}

	p := self.logQueue[self.head]
	self.logQueue[self.head] = nil
	self.head = (self.head + 1) % self.cap
	self.queue_len--

	return p
}

/*读取队列信息*/
func (self *LogQueue) GetQueueInfo() (len, cap int) {
	return self.queue_len, self.cap
}

/*清空队列*/
func (self *LogQueue) Clear() {

	for i := 0; i < self.cap; i++ {
		self.logQueue[i] = nil
	}
	self.head, self.rear, self.queue_len = 0, 0, 0

}

/*打印队列内部信息队列信息*/
func (self *LogQueue) PrintQueueInfo() {

	fmt.Printf("Queue:Capacity=%d, Len=%d, Head=%d,Read=%d\n", self.cap, self.queue_len, self.head, self.rear)

}
