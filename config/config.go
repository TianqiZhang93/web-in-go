package conf

import (
	"fmt"
	"reflect"

	"github.com/go-gcfg/gcfg"
)

type APIConfItem interface {
	Check() error
}

// Each item of APIConf implements APIConfItem
type APIConf struct {
	Server ServerConfig
}

// Load config and do check.
func Load(confPath string) (*APIConf, error) {
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

func load(confPath string) (*APIConf, error) {
	conf := APIConf{}
	err := gcfg.ReadFileInto(&conf, confPath)

	return &conf, err
}

func (a *APIConf) check() error {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		if err := v.Field(i).Addr().Interface().(APIConfItem).Check(); err != nil {
			return fmt.Errorf("%v Check(): %v", t.Field(i).Name, err)
		}
	}

	return nil
}
