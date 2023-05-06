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
	log.Info("blade begin stop")
	atomic.StoreInt32(&IsMinerNotRunning, 1)
}

//等待chan
func CloseSync(pos string) {
	if isClose() {
		log.Info("blade before stop channel", pos, 1)
		<-MyChan
		log.Info("blade after stop channel", pos, 1)
	}
}
func PassCloseChan(pos string) {
	defer func() {
		if r := recover(); r != nil {
			log.Info("blade recover PassCloseChan", pos, 1)
		}
	}()
	if isClose() {
		close(MyChan)
		log.Info("blade  PassCloseChan", pos, 1)
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
	log.Info("blade before MainCloseChan sleep")
	time.Sleep(10 * time.Second)
	close(MyChan)
	log.Info("blade MainCloseChan")
}
