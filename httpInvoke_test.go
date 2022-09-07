package fss

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fss/eumTaskType"
	"testing"
	"time"
)

func Test_httpPull(t *testing.T) {
	fs.Initialize[Module]("fss test")
	configure.SetDefault("FSS.Server", "http://127.0.0.1:888")
	lst := httpPull(1)
	for _, vo := range lst.ToArray() {
		httpInvoke(invokeRequest{
			TaskGroupId:  vo.TaskGroupId,
			NextTimespan: 0,
			Progress:     0,
			Status:       eumTaskType.Working,
			RunSpeed:     10,
			Log:          LogRequest{},
			Data:         vo.Data,
		})
		time.Sleep(100 * time.Millisecond)

		httpInvoke(invokeRequest{
			TaskGroupId:  vo.TaskGroupId,
			NextTimespan: 0,
			Progress:     0,
			Status:       eumTaskType.Success,
			RunSpeed:     10,
			Log:          LogRequest{},
			Data:         vo.Data,
		})
	}
}
