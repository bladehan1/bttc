package myconst

import (
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

// 开始关闭
var IsMinerNotRunning int32 = 0

//等待chan
var MyChan = make(chan int)

func CloseSignal() {
	atomic.StoreInt32(&IsMinerNotRunning, 1)
}

//等待chan
func CloseSync() {
	if isClose() {
		log.Info("blade before stop channel")
		<-MyChan
		log.Info("blade after stop channel")
	}
}
func PassCloseChan(pos string) {
	defer func() {
		if r := recover(); r != nil {
			log.Info("blade recover PassCloseChan", pos)
		}
	}()
	if isClose() {
		close(MyChan)
		log.Info("blade  PassCloseChan", pos)
	}
}

func isClose() bool {
	return atomic.LoadInt32(&IsMinerNotRunning) == 1
}

func MainCloseChan() {
	defer func() {
		if r := recover(); r != nil {
			log.Info("blade recover MainCloseChan")
		}
	}()
	time.Sleep(time.Second)
	close(MyChan)
	log.Info("blade MainCloseChan")
}
