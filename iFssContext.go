package fss

import (
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/stopwatch"
	"github.com/farseer-go/fss/eumTaskType"
	"time"
)

type IFssContext interface {
	// SetProgress 上传进度0-100
	SetProgress(rate int)
	// SetNextAt 本次执行完后，下一次执行的间隔时间
	SetNextAt(d time.Duration)
	// Logger 写入到FSS平台的日志
	Logger(logLevel eumLogLevel.Enum, log string)
}

type fssContext struct {
	watch        *stopwatch.Watch // 任务执行耗时
	task         taskVO           // 任务
	isDebug      bool             // 是否调试状态
	taskStatus   eumTaskType.Enum // 当前任务状态
	nextTimespan int64            // 下一次执行时间
	progress     int              // 任务执行进度：0-100
}

func newContext(sw *stopwatch.Watch, task taskVO) fssContext {
	context := fssContext{
		watch:      sw,
		task:       task,
		taskStatus: eumTaskType.Working,
	}
	if context.task.Data == nil {
		context.task.Data = make(map[string]string)
	}
	return context
}

// success 成功后执行
func (r *fssContext) success() {
	r.taskStatus = eumTaskType.Success
	if !r.isDebug {
		httpInvoke(invokeRequest{
			TaskGroupId:  r.task.TaskGroupId,
			NextTimespan: r.nextTimespan,
			Progress:     100,
			Status:       r.taskStatus,
			RunSpeed:     r.watch.ElapsedMilliseconds(),
			Log:          LogRequest{},
			Data:         r.task.Data,
		})
	}
}

// fail 执行失败
func (r *fssContext) fail(logContent string) {
	r.taskStatus = eumTaskType.Fail
	if !r.isDebug {
		var log LogRequest
		if logContent != "" {
			log = LogRequest{
				LogLevel: eumLogLevel.Error,
				Log:      logContent,
				CreateAt: time.Now(),
			}
		}
		httpInvoke(invokeRequest{
			TaskGroupId:  r.task.TaskGroupId,
			NextTimespan: r.nextTimespan,
			Progress:     r.progress,
			Status:       r.taskStatus,
			RunSpeed:     r.watch.ElapsedMilliseconds(),
			Log:          log,
			Data:         r.task.Data,
		})
	}
}

// activateTask 激活任务
func (r *fssContext) activateTask() {
	if !r.isDebug && r.taskStatus == eumTaskType.Working {
		httpInvoke(invokeRequest{
			TaskGroupId:  r.task.TaskGroupId,
			NextTimespan: r.nextTimespan,
			Progress:     r.progress,
			Status:       r.taskStatus,
			RunSpeed:     r.watch.ElapsedMilliseconds(),
			Log:          LogRequest{},
			Data:         r.task.Data,
		})
	}
}

// SetProgress 上传进度0-100
func (r *fssContext) SetProgress(rate int) {
	if rate < 1 || rate > 100 {
		return
	}
	r.progress = rate
}

// SetNextAt 本次执行完后，下一次执行的间隔时间
func (r *fssContext) SetNextAt(d time.Duration) {
	r.nextTimespan = time.Now().Add(d).UnixMilli()
}

// Logger 写入到FSS平台的日志
func (r *fssContext) Logger(logLevel eumLogLevel.Enum, log string) {
	flog.Log(logLevel, log)
	if !r.isDebug {
		httpInvoke(invokeRequest{
			TaskGroupId:  r.task.TaskGroupId,
			NextTimespan: r.nextTimespan,
			Progress:     r.progress,
			Status:       r.taskStatus,
			RunSpeed:     r.watch.ElapsedMilliseconds(),
			Log: LogRequest{
				LogLevel: logLevel,
				Log:      log,
				CreateAt: time.Now(),
			},
			Data: r.task.Data,
		})
	}
}

// GetData 获取数据
func (r *fssContext) GetData(key string) string {
	return r.task.Data[key]
}

// SetData 获取数据
func (r *fssContext) SetData(key string, val string) {
	r.task.Data[key] = val
}
