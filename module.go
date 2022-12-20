package fss

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/fs/snowflake"
	"os"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	return nil
}

func (module Module) PreInitialize() {
	hostname, _ := os.Hostname()
	defaultClient = clientVO{
		ClientId:        snowflake.GenerateId(),
		ClientIp:        fs.AppIp,
		ClientName:      hostname,
		ClientJobs:      collections.NewDictionary[string, jobFunc](),
		WorkFinishEvent: make(chan int, 100),
	}
}

func (module Module) Initialize() {
}

func (module Module) PostInitialize() {
	// 后台拉取任务
	go timingPullTask()
}

func (module Module) Shutdown() {
}
