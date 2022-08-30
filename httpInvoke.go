package fss

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fss/eumTaskType"
	"github.com/farseer-go/utils/http"
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

// 从fss服务端拉取任务
func httpPull() collections.List[taskVO] {
	url := http.AddHttpPrefix(collections.NewList(getServerConfig()...).Rand()) + "/task/pull"
	postData := map[string]any{"TaskCount": getPullCountConfig()}
	httpHead := client.getHttpHead()
	var rsp = http.PostJson[core.ApiResponse[collections.List[taskVO]]](url, httpHead, postData, 500)
	if !rsp.Status {
		return collections.NewList[taskVO]()
	}
	return rsp.Data
}

// 向fss服务端提交任务报告
func httpInvoke(request invokeRequest) bool {
	url := http.AddHttpPrefix(collections.NewList(getServerConfig()...).Rand()) + "/task/JobInvoke"
	httpHead := client.getHttpHead()
	var rsp = http.PostJson[core.ApiResponseString](url, httpHead, request, 500)
	return rsp.Status
}
