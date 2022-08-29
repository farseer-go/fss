package fss

import "github.com/farseer-go/fs/configure"

type fssConfig struct {
	Server    string // fss服务端地址
	PullCount int    // 客户端每次拉取数量
	WorkCount int    // 允许最大并行执行的任务数量
}

// 允许最大并行执行的任务数量
func getWorkCountConfig() int {
	return configure.GetInt("FSS.WorkCount")
}

// 客户端每次拉取数量
func getPullCountConfig() int {
	return configure.GetInt("FSS.PullCount")
}

// fss服务端地址
func getServerConfig() []string {
	return configure.GetStrings("FSS.Server")
}
