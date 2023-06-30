package config

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck_OK(t *testing.T) {
	confPath := "./testdata/server.yaml"

	conf, err := Load(confPath)
	assert.NoError(t, err)

	expect := &LocalConfig{
		Server: ServerConfig{
			ServicePort:   8001,
			MonitorPort:   8002,
			PprofPort:     8003,
			WhiteList:     map[string][]string{"test": {"127.0.0.1", "::1"}},
			AccessLogPath: "../log/access.log",
		},
	}
	assert.True(t, reflect.DeepEqual(conf, expect))
}

func TestCheck_EmptyWhitelist(t *testing.T) {
	confPath := "./testdata/server_invalid_port.yaml"

	_, err := Load(confPath)
	fmt.Println(err)
	assert.True(t, strings.Contains(fmt.Sprintf("%v", err), "should > 0"))
}
