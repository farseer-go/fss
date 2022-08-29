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
}

func (module Module) Initialize() {
}

func (module Module) PostInitialize() {
	hostname, _ := os.Hostname()
	client = clientVO{
		Id:   snowflake.GenerateId(),
		Ip:   fs.AppIp,
		Name: hostname,
		Jobs: collections.NewList[string](),
	}
	httpHead = map[string]any{
		"ClientIp":   client.Ip,
		"ClientId":   client.Id,
		"ClientName": client.Name,
		"ClientJobs": "",
	}
}

func (module Module) Shutdown() {
}
