package fss

import (
	"time"
)

// 定时拉取任务
func timingPullTask() {
	// 等待500ms，等待job注册进来
	time.Sleep(500 * time.Millisecond)
	// 第一次时不需要等2秒，应立即拉取任务
	client.WorkFinishEvent <- 0
	for {
		select {
		//case <-client.WorkFinishEvent:
		//	pullTask()
		case <-time.After(time.Duration(getPullIntervalConfig()) * time.Millisecond):
			pullTask()
		}
	}
}

// 拉取任务
func pullTask() {
	pullCount := getPullCountConfig() - client.WaitCount - client.WorkCount
	if pullCount > 0 {
		lstTask := httpPull(pullCount)
		// 拉取到任务后，将任务发到chan队列中
		for _, task := range lstTask.ToArray() {
			if task.TaskGroupId > 0 {
				go executeTask(task)
			}
		}
	}
}
