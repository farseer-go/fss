package fss

import (
	"fmt"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/stopwatch"
	"time"
)

// 拉取任务
func pullTask() {
	// 等待500ms，等待job注册进来
	time.Sleep(500 * time.Millisecond)
	for {
		lstTask := httpPull()
		// 拉取到任务后，将任务发到chan队列中
		for _, task := range lstTask.ToArray() {
			client.taskQueue <- task
		}
	}
}

func workTask() {
	// 遍历通道中的任务
	for task := range client.taskQueue {
		sw := stopwatch.StartNew()
		receiveContext := newContext(sw, task)
		// 如果拉回来的任务，本地不支持，则停止
		if !client.ClientJobs.ContainsKey(task.JobName) {
			message := fmt.Sprintf("未找到任务实现类：任务组：TaskGroupId=%d，Caption=%s，JobName=%s", task.TaskGroupId, task.Caption, task.JobName)
			flog.Error(message)
			receiveContext.Fail(message)
			continue
		}
	}
}
