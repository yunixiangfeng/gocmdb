package plugins

import (
	"time"

	"gocmdb/agent/config"
)

type CyclePlugin interface {
	Name() string
	Init(*config.Config)
	NextTime() time.Time
	Call() (interface{}, error)
	Pipline() chan interface{}
}

type TaskPlugin interface {
	Name() string
	Init(*config.Config)
	Call(params string) (interface{}, error)
}
