package scheduler

import (
	"github.com/Galagoshin/GoLogger/logger"
	"time"
)

type UnixTimeTask struct {
	keys map[any]time.Time
}

func (task *UnixTimeTask) RemoveKey(id any) {
	delete(task.keys, id)
}

func (task *UnixTimeTask) SetKey(id any, duration time.Duration) {
	if task.keys == nil {
		task.keys = make(map[any]time.Time)
	}
	task.keys[id] = time.Now().Add(duration)
}

func (task *UnixTimeTask) ExistsKey(id any) bool {
	_, exists := task.keys[id]
	return exists
}

func (task *UnixTimeTask) GetTimeLeft(id any) (time.Duration, bool) {
	key, exists := task.keys[id]
	if !exists {
		return time.Duration(0), exists
	}
	return time.Duration(key.Unix() - time.Now().Unix()), exists
}

func (task *UnixTimeTask) IsComplete(id any) bool {
	val, exists := task.GetTimeLeft(id)
	if exists {
		return val <= 0
	} else {
		logger.Warning("Key id not found.")
		return false
	}
}
