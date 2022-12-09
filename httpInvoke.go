package fss

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/stopwatch"
	"github.com/farseer-go/fss/eumTaskType"
	"github.com/farseer-go/utils/http"
	"time"
)

type invokeRequest struct {
	TaskGroupId  int                                    // 任务组ID
	NextTimespan int64                                  // 下次执行时间
	Progress     int                                    // 当前进度
	Status       eumTaskType.Enum                       // 执行状态
	RunSpeed     int64                                  // 执行速度
	Log          LogRequest                             // 日志
	Data         collections.Dictionary[string, string] // 数据
}

// LogRequest 日志
type LogRequest struct {
	LogLevel eumLogLevel.Enum
	Log      string
	CreateAt time.Time
}

// 从fss服务端拉取任务
func httpPull(taskCount int) collections.List[taskVO] {
	url := http.AddHttpPrefix(collections.NewList(getServerConfig()...).Rand()) + "/task/pull"
	sw := stopwatch.StartNew()
	postData := map[string]any{"TaskCount": taskCount}
	httpHead := client.getHttpHead()

	var rsp core.ApiResponse[collections.List[taskVO]]
	http.NewClient(url).Head(httpHead).Body(postData).PostUnmarshal(&rsp)
	defer func() {
		flog.ComponentInfof("fss", "本次拉取任务%d条，耗时：%s", rsp.Data.Count(), sw.GetMillisecondsText())
	}()
	if !rsp.Status {
		flog.Warning("%s request Warning:%s", url, rsp.StatusMessage)
		return collections.NewList[taskVO]()
	}
	return rsp.Data
}

// 向fss服务端提交任务报告
func httpInvoke(request invokeRequest) bool {
	url := http.AddHttpPrefix(collections.NewList(getServerConfig()...).Rand()) + "/task/jobinvoke"
	httpHead := client.getHttpHead()
	var rsp core.ApiResponseString
	http.NewClient(url).Head(httpHead).Body(request).PostUnmarshal(&rsp)

	if !rsp.Status {
		flog.Warning("%s request Warning:%s", url, rsp.StatusMessage)
	}
	return rsp.Status
}
