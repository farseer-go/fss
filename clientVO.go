package fss

import "github.com/farseer-go/collections"

var client clientVO

type clientVO struct {
	ClientId   int64                                   // 客户端ID
	ClientIp   string                                  // 客户端IP
	ClientName string                                  // 客户端名称
	ClientJobs collections.Dictionary[string, jobFunc] // 客户端能执行的任务
	taskQueue  chan taskVO
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
