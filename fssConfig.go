package fss

type fssConfig struct {
	Server    string // fss服务端地址
	PullCount int    // 客户端每次拉取数量
	WorkCount int    // 允许最大并行执行的任务数量
}
