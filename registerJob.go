package fss

type jobFunc func(context IFssContext) bool

// RegisterJob 注册JOB
func RegisterJob(jobName string, jobFn jobFunc) {
	if client.WorkFinishEvent == nil {
		panic("使用fss组件时，需依赖fss.Module模块")
	}
	if client.ClientJobs.ContainsKey(jobName) {
		panic("jobName：" + jobName + "，已存在")
	}
	client.ClientJobs.Add(jobName, jobFn)
}
