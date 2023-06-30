package conf

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	confPath := "./testdata/server.conf"

	conf, err := Load(confPath)
	assert.NoError(t, err)

	expect := &APIConf{
		Server: ServerConfig{8001, 8002, 8003, []string{"127.0.0.1", "::1"}},
	}
	assert.True(t, reflect.DeepEqual(conf, expect))
}

func TestCheck_EmptyWhitelist(t *testing.T) {
	confPath := "./testdata/server_emptyWhitelist.conf"

	_, err := Load(confPath)
	assert.True(t, strings.Contains(fmt.Sprintf("%v", err), "no item in WhiteList"))
}
