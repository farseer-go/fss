package fss

import (
	"github.com/farseer-go/fs"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestModule_PreInitialize(t *testing.T) {
	fs.Initialize[Module]("fss test")
	assert.Equal(t, 3, strings.Count(client.ClientIp, "."))
	assert.Less(t, int64(0), client.ClientId)
	assert.NotEmpty(t, client.ClientName)
	assert.NotNil(t, client.ClientJobs)
	assert.NotNil(t, client.WorkFinishEvent)
}
