package test

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fss"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestModule_PreInitialize(t *testing.T) {
	fs.Initialize[fss.Module]("fss test")
	assert.Equal(t, 3, strings.Count(fss.GetClient().ClientIp, "."))
	assert.Less(t, int64(0), fss.GetClient().ClientId)
	assert.NotEmpty(t, fss.GetClient().ClientName)
	assert.NotNil(t, fss.GetClient().ClientJobs)
	assert.NotNil(t, fss.GetClient().WorkFinishEvent)
}
