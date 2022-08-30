package fss

type jobFunc func(context IFssContext)

// RegisterJob 注册JOB
func RegisterJob(jobName string, jobFn jobFunc) {
	if client.ClientJobs.ContainsKey(jobName) {
		panic("jobName：" + jobName + "，已存在")
	}
	client.ClientJobs.Add(jobName, jobFn)
}
