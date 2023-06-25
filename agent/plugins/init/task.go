package init

import (
	"gocmdb/agent/plugins"
	"gocmdb/agent/plugins/task"
)

func init() {
	plugins.DefaultManager.RegisterTask(&task.Process{})
}
