package fss

import (
	"github.com/farseer-go/collections"
	"time"
)

// fss服务端拉取回来的任务
type taskVO struct {
	TaskGroupId int                                    // 任务组ID
	Caption     string                                 // 任务组标题
	JobName     string                                 // 实现Job的特性名称（客户端识别哪个实现类）
	StartAt     time.Time                              // 开始时间
	RunAt       time.Time                              // 实际执行时间
	RunSpeed    int                                    // 运行耗时
	ClientHost  string                                 // 客户端
	ClientIp    string                                 // 客户端IP
	ClientName  string                                 // 客户端名称
	Progress    int                                    // 进度0-100
	CreateAt    time.Time                              // 任务创建时间
	SchedulerAt time.Time                              // 调度时间
	Data        collections.Dictionary[string, string] // 本次执行任务时的Data数据
}
