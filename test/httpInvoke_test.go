package test

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fss"
	"github.com/farseer-go/fss/eumTaskType"
	"testing"
	"time"
)

func Test_httpPull(t *testing.T) {
	fs.Initialize[fss.Module]("fss test")
	configure.SetDefault("FSS.Server", "http://127.0.0.1:888")
	lst := fss.httpPull(1)
	for _, vo := range lst.ToArray() {
		fss.httpInvoke(fss.invokeRequest{
			TaskGroupId:  vo.TaskGroupId,
			NextTimespan: 0,
			Progress:     0,
			Status:       eumTaskType.Working,
			RunSpeed:     10,
			Log:          fss.LogRequest{},
			Data:         vo.Data,
		})
		time.Sleep(100 * time.Millisecond)

		fss.httpInvoke(fss.invokeRequest{
			TaskGroupId:  vo.TaskGroupId,
			NextTimespan: 0,
			Progress:     0,
			Status:       eumTaskType.Success,
			RunSpeed:     10,
			Log:          fss.LogRequest{},
			Data:         vo.Data,
		})
	}
}
