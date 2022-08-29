package fss

type jobFunc func(context IFssContext)

// RegisterJob 注册JOB
func RegisterJob(jobName string, jobFn jobFunc) {

}
