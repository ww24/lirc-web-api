package service

import (
	"time"

	"github.com/ww24/lirc-web-api/datastore"
)

const (
	onceScheduleTaskPrefix    = "task:once:"
	routineScheduleTaskPrefix = "task:routine:"
)

var store = datastore.NewLevelDBStore()

// Scheduler interface
type Scheduler interface {
	Execute()
	Register()
	Unregister()
}

// Task structure
type Task struct {
	ID        string    `codec:"-" json:"id"`
	Signals   []*Signal `codec:"signals" json:"signals"`
	CreatedAt time.Time `codec:"created_at" json:"created_at"`
}

// Execute implement Scheduler
func (t *Task) Execute() error {
	for _, signal := range t.Signals {
		err := SendSignal(&SendSignalParam{Signal: signal})
		if err != nil {
			return err
		}
	}
	return nil
}

// OnceScheduleTask structure
type OnceScheduleTask struct {
	*Task
	Time time.Time `json:"time"`
}

// Register implement Scheduler
func (t *OnceScheduleTask) Register() error {
	prefix := onceScheduleTaskPrefix
	return nil
}

// Unregister implement Scheduler
func (t *OnceScheduleTask) Unregister() error {
	prefix := onceScheduleTaskPrefix
	return nil
}

// RoutineScheduleTask structure
type RoutineScheduleTask struct {
	*Task
	Rule string `codec:"rule" json:"rule"`
}

// Register implement Scheduler
func (t *RoutineScheduleTask) Register() error {
	prefix := routineScheduleTaskPrefix
	return nil
}

// Unregister implement Scheduler
func (t *RoutineScheduleTask) Unregister() error {
	prefix := routineScheduleTaskPrefix
	return nil
}

// AddScheduler .
func AddScheduler(s Scheduler) {
	s.Register()
}

// RemoveScheduler .
func RemoveScheduler(s Scheduler) {

}
