package system

import (
	"btquan/common"
	"btquan/x/mlog"
	"btquan/x/timer"
	"fmt"
	"time"
)

var logSys = mlog.NewTagLog("System")

func Start() (tkWorker *VideoWorker) {
	tkWorker = newCacheVideoWorker()
	return
}

func (tkWorker *VideoWorker) Launch() {
	go tkWorker.startCache()
}

func (c *VideoWorker) startCache() {
	every15Minute := time.Tick(time.Duration(common.ConfigSystemBooking.CyclePushDay) * time.Minute)
	//every2Minute := time.Tick(time.Duration(common.ConfigSystemBooking.CyclePushVideo) * time.Minute)
	daily := timer.NewDailyTimer(common.ConfigSystemBooking.TimeSetCache, 0)
	daily.Start()
	for {
		select {
		case <-every15Minute:

		case action := <-c.VideoUpdate:
			c.VideoWorking(action)
		case <-daily.C:
			fmt.Println("======== EVERYDAY ==========")
			c.removeTksVideoWorkerDay()
		}

	}
}
