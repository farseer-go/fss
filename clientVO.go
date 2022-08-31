package fss

import "github.com/farseer-go/collections"

var client clientVO

type clientVO struct {
	ClientId        int64                                   // 客户端ID
	ClientIp        string                                  // 客户端IP
	ClientName      string                                  // 客户端名称
	ClientJobs      collections.Dictionary[string, jobFunc] // 客户端能执行的任务
	WorkCount       int                                     // 正在执行
	WaitCount       int                                     // 等待中的任务数
	WorkFinishEvent chan int                                // 任务完成后通知
}

// 转换成http head
func (receiver clientVO) getHttpHead() map[string]any {
	return map[string]any{
		"ClientIp":   client.ClientIp,
		"ClientId":   client.ClientId,
		"ClientName": client.ClientName,
		"ClientJobs": client.ClientJobs.Keys().ToString(","),
	}
}
