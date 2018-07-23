package system

import (
	"btquan/x/mrw/event"
)

type VideoWorker struct {
	VideoUpdate chan *VideoAction
	doneAction  *event.Hub
}

func (tw *VideoWorker) TriggerVideoAction(action *VideoAction) {
	tw.VideoUpdate <- action
}

func newCacheVideoWorker() *VideoWorker {
	return &VideoWorker{
		VideoCaches: make(map[string]*video.Video, 0),
		VideoUpdate: make(chan *VideoAction, event.MediumHub),
	}
}

func (tkDay *VideoWorker) GetVideoByID(idVideo string) (*video.Video, error) {
	if tk, ok := tkDay.VideoCaches[idVideo]; ok {
		return tk, nil
	}
	return nil, nil
}

func (tw *VideoWorker) OnActionDone() (event.Line, event.Cancel) {
	return tw.doneAction.NewLine()
}

func (tkDay *VideoWorker) removeTksVideoWorkerDay() {
	if tkDay == nil {
		if len(tkDay.VideoCaches) > 0 {
			for k := range tkDay.VideoCaches {
				delete(tkDay.VideoCaches, k)
			}
		}
	}
}
