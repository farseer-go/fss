package fss

import (
	"context"
	"fmt"
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/stopwatch"
	"github.com/farseer-go/tasks"
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
		case <-client.WorkFinishEvent:
			pullTask()
		case <-time.After(2 * time.Second):
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
			go executeTask(task)
		}
	}
}

// 执行任务
func executeTask(task taskVO) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	sw := stopwatch.New()
	receiveContext := newContext(sw, task)
	// 如果拉回来的任务，本地不支持，则停止
	if !client.ClientJobs.ContainsKey(task.JobName) {
		message := fmt.Sprintf("未找到任务实现类：任务组：TaskGroupId=%d，Caption=%s，JobName=%s", task.TaskGroupId, task.Caption, task.JobName)
		flog.Error(message)
		receiveContext.fail(message)
		return
	}

	client.WaitCount++
	// 任务执行客户端
	fssJob := client.ClientJobs.GetValue(task.JobName)

	// 计划时间还没到
	waitTimeSpan := task.StartAt.Sub(time.Now())
	if waitTimeSpan > 0 {
		time.Sleep(waitTimeSpan)
	}
	// 等待任务减1，工作中+1
	client.WaitCount--
	client.WorkCount++
	defer func() {
		client.WorkCount--
	}()

	// 定时激活任务
	tasks.RunNow("fss.activateTask", 3*time.Second, func(context *tasks.TaskContext) { receiveContext.activateTask() }, ctx)
	sw.Start()
	var result bool

	// 执行任务
	try := exception.Try(func() {
		result = fssJob(&receiveContext)
	})
	try.CatchException(func(exp any) {
		flog.Errorf("taskGroupId=%d,caption=%s：%s", receiveContext.task.TaskGroupId, receiveContext.task.Caption, exp)
		receiveContext.fail(exp.(string))
	})

	cancelFunc()
	if result {
		receiveContext.success()
	} else {
		receiveContext.fail("")
	}
}
