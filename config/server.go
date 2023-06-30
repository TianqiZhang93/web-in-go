package config

import "fmt"

type ServerConfig struct {
	ServicePort int `json:"service_port" yaml:"service_port"`
	MonitorPort int `json:"monitor_port" yaml:"monitor_port"`
	PprofPort   int `json:"pprof_port" yaml:"pprof_port"`

	WhiteList map[string][]string `json:"whitelist" yaml:"whitelist"`

	AccessLogPath string `json:"access_log_path" yaml:"access_log_path"`
}

// implement APIConfItem
func (s *ServerConfig) Check() error {
	if s == nil {
		return fmt.Errorf("nil data")
	}

	if s.ServicePort <= 0 || s.MonitorPort <= 0 || s.PprofPort <= 0 {
		return fmt.Errorf("service_port[%d] monitor_port[%d] pprof_port[%d] should > 0",
			s.ServicePort, s.MonitorPort, s.PprofPort)
	}

	return nil
}
