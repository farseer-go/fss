package fss

import (
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fss/eumTaskType"
	"time"
)

type invokeRequest struct {
	TaskGroupId  int               // 任务组ID
	NextTimespan int64             // 下次执行时间
	Progress     int               // 当前进度
	Status       eumTaskType.Enum  // 执行状态
	RunSpeed     int64             // 执行速度
	Log          LogRequest        // 日志
	Data         map[string]string // 数据
}

// LogRequest 日志
type LogRequest struct {
	LogLevel eumLogLevel.Enum
	Log      string
	CreateAt time.Time
}

func httpInvoke(request invokeRequest) {

}
