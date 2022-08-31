package fss

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/utils/snowflake"
	"math/rand"
	"os"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	return []modules.FarseerModule{modules.FarseerKernelModule{}}
}

func (module Module) PreInitialize() {
	hostname, _ := os.Hostname()
	snowflake.Init(parse.HashCode64(hostname), rand.Int63n(32))
	client = clientVO{
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
