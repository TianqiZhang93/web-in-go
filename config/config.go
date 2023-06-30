package config

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"gopkg.in/yaml.v2"
)

type ConfChecker interface {
	Check() error
}

type LocalConfig struct {
	Server ServerConfig `json:"server" yaml:"server"`
}

// Load config and do check.
func Load(confPath string) (*LocalConfig, error) {
	conf, err := load(confPath)
	if err != nil {
		return conf, err
	}

	err = conf.check()
	if err != nil {
		return conf, err
	}

	return conf, nil
}
func load(confPath string) (*LocalConfig, error) {
	fileData, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile(%s): %s", confPath, err.Error())
	}

	conf := LocalConfig{}
	err = yaml.Unmarshal(fileData, &conf)
	if err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal(): %s", err.Error())
	}

	return &conf, err
}

func (a *LocalConfig) check() error {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		if err := v.Field(i).Addr().Interface().(ConfChecker).Check(); err != nil {
			return fmt.Errorf("%v Check(): %v", t.Field(i).Name, err)
		}
	}

	return nil
}
