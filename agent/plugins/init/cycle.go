package init

import (
	"gocmdb/agent/plugins"
	"gocmdb/agent/plugins/cycle"
)

func init() {
	plugins.DefaultManager.RegisterCycle(&cycle.Heartbeat{})
	plugins.DefaultManager.RegisterCycle(&cycle.Register{})
	plugins.DefaultManager.RegisterCycle(&cycle.Resource{})
}
