package system

import (
	"errors"
)

func (action *VideoAction) handlerAction(Video *video.Video) {
	switch action.Action {

	default:
		err := errors.New("No Action")
		action.SetError(err)
	}
}
