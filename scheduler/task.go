package scheduler

type TaskType int64

type SchedulingTask interface {
	Destroy()
	Run(...any)
}
