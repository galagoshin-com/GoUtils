package scheduler

import (
	"github.com/galagoshin-com/GoLogger/logger"
	"time"
)

type OneTimeTask struct {
	Duration   time.Duration
	OnComplete func(...any)
	timer      *time.Timer
}

func (task *OneTimeTask) Destroy() {
	<-task.timer.C
	task.timer.Stop()
}

func (task *OneTimeTask) Run(args ...any) {
	task.timer = time.NewTimer(task.Duration)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Error(r.(error))
			}
		}()
		<-task.timer.C
		task.OnComplete(args...)
		task.timer.Reset(task.Duration)

	}()
}
