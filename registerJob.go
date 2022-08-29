package fss

type jobFunc func(context IFssContext)

func RegisterJob(jobName string, jobFn jobFunc) {

}
