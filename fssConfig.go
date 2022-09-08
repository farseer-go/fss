package fss

import (
	"github.com/farseer-go/fs/configure"
	"runtime"
)

type fssConfig struct {
	Server    string // fss服务端地址
	PullCount int    // 客户端每次拉取数量
	WorkCount int    // 允许最大并行执行的任务数量
}

// 允许最大并行执行的任务数量
func getWorkCountConfig() int {
	workCount := configure.GetInt("FSS.WorkCount")
	if workCount == 0 {
		workCount = runtime.NumCPU()
	}
	return workCount
}

// 客户端每次拉取数量
func getPullCountConfig() int {
	pullCount := configure.GetInt("FSS.PullCount")
	if pullCount == 0 {
		pullCount = runtime.NumCPU()
	}
	return pullCount
}

// fss服务端地址
func getServerConfig() []string {
	return configure.GetStrings("FSS.Server")
}
