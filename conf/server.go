// config.go - config for api

/*
modification history
--------------------
2020/05/15, by NKztq, create
*/
package conf

import "fmt"

type ServerConfig struct {
	ServPort    int
	MonitorPort int
	PprofPort   int

	WhiteList []string
}

// implement APIConfItem
func (s *ServerConfig) Check() error {
	if len(s.WhiteList) == 0 {
		return fmt.Errorf("no item in WhiteList")
	}

	return nil
}
