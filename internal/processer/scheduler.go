package processer

import (
	"github.com/ThreatCode/Proxy-Check/internal"
	"github.com/ThreatCode/Proxy-Check/internal/model"
)

func ScheduleTask(task *model.Task, taskQueue chan *model.Task) {
	taskQueue <- task
	internal.State.TaskScheduled()
}
