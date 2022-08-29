package fss

import "github.com/farseer-go/collections"

var client clientVO

type clientVO struct {
	Id   int64                    // 客户端ID
	Ip   string                   // 客户端IP
	Name string                   // 客户端名称
	Jobs collections.List[string] // 客户端能执行的任务
}
