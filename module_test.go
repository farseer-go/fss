package fss

import (
	"github.com/farseer-go/fs"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestModule_PreInitialize(t *testing.T) {
	fs.Initialize[Module]("fss test")
	assert.Equal(t, 4, strings.Count(client.ClientIp, "."))
}
