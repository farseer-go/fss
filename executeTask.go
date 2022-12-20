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

// 执行任务
func executeTask(task taskVO) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	sw := stopwatch.New()
	receiveContext := newContext(sw, task)
	// 如果拉回来的任务，本地不支持，则停止
	if !defaultClient.ClientJobs.ContainsKey(task.JobName) {
		message := fmt.Sprintf("未找到任务实现类：任务组：TaskGroupId=%d，Caption=%s，JobName=%s", task.TaskGroupId, task.Caption, task.JobName)
		flog.Error(message)
		receiveContext.fail(message)
		return
	}

	defaultClient.WaitCount++
	// 任务执行客户端
	fssJob := defaultClient.ClientJobs.GetValue(task.JobName)

	// 计划时间还没到
	waitTimeSpan := task.StartAt.Sub(time.Now())
	if waitTimeSpan > 0 {
		flog.ComponentInfof("fss", "（%d）任务组：%s，计划时间还没到，休眠%d ms", task.TaskGroupId, task.Caption, waitTimeSpan.Milliseconds())
		time.Sleep(waitTimeSpan)
	} else {
		flog.ComponentInfof("fss", "（%d）任务组：%s，延迟了%d ms", task.TaskGroupId, task.Caption, waitTimeSpan.Milliseconds())
	}

	// 等待任务减1，工作中+1
	defaultClient.WaitCount--
	defaultClient.WorkCount++
	defer func() {
		defaultClient.WorkCount--
	}()

	// 定时激活任务
	tasks.RunNow("fss.activateTask", 3*time.Second, func(context *tasks.TaskContext) { receiveContext.activateTask() }, ctx)
	sw.Start()
	var result bool

	// 执行任务
	try := exception.Try(func() {
		sw := stopwatch.StartNew()
		result = fssJob(&receiveContext)
		flog.ComponentInfof("fss", "%s，耗时：%s", task.JobName, sw.GetMillisecondsText())
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
