package scheduler

import (
	"github.com/Galagoshin/GoLogger/logger"
	"time"
)

type RepeatingTask struct {
	Duration   time.Duration
	OnComplete func(...any)
	timer      *time.Timer
}

func (task *RepeatingTask) Destroy() {
	<-task.timer.C
	task.timer.Stop()
}

func (task *RepeatingTask) Run(args ...any) {
	task.timer = time.NewTimer(task.Duration)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Error(r.(error))
			}
		}()
		for {
			<-task.timer.C
			task.OnComplete(args...)
			task.timer.Reset(task.Duration)
		}
	}()
}
