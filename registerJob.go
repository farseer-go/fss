package fss

import "github.com/farseer-go/fs/flog"

type jobFunc func(context IFssContext) bool

// RegisterJob 注册JOB
func RegisterJob(jobName string, jobFn jobFunc) {
	if defaultClient.WorkFinishEvent == nil {
		panic("使用fss组件时，需依赖fss.Module模块")
	}
	if defaultClient.ClientJobs.ContainsKey(jobName) {
		panic("jobName：" + jobName + "，已存在")
	}
	defaultClient.ClientJobs.Add(jobName, jobFn)

	flog.ComponentInfof("fss", "注册任务：%s", jobName)
}
