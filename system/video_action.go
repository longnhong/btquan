package system

import (
	"btquan/x/mlog"
	"btquan/x/mrw/encode"
)

var logAction = mlog.NewTagLog("system")

type VideoAction struct {
	VideoID   string
	CusID     string
	PushToken string
	Extra     encode.RawMessage `json:"extra"`
	Error     VideoActionError  `json:"error"`
	doneC     chan struct{}
	used      bool // must be trigger at most once
}
type VideoActionError struct {
	s string
}

func (a *VideoAction) Done() bool {
	a.doneC <- struct{}{}
	return a.GetError() == nil
}

func (e *VideoActionError) Error() string {
	return e.s
}

func (e *VideoActionError) GetError() error {
	if len(e.s) > 0 {
		return e
	}
	return nil
}

func (a *VideoAction) SetError(err error) {
	if err == nil {
		return
	}
	logAction.Errorf("SetError", err)
	a.Error = VideoActionError{s: err.Error()}
}

func (a *VideoAction) GetError() error {
	return a.Error.GetError()
}

func (a *VideoAction) Wait() (*video.Video, error) {
	if a.doneC == nil {
		logAction.Errorf("Wait()", "no done channel")
		panic("no done channel")
	}
	<-a.doneC
	return a.Video, a.GetError()
}

func NewVideoAction() *VideoAction {
	var a = &VideoAction{
		doneC: make(chan struct{}, 1),
	}
	return a
}
