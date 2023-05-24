package ip2location

import (
	"github.com/fimreal/rack/module"
	"github.com/spf13/cobra"
)

const (
	ID            = "ip2location"
	Comment       = "ip2location sdk"
	RoutePrefix   = "/ip2location"
	DefaultEnable = false
)

var Module = module.Module{
	ID:      ID,
	Comment: Comment,
	// gin route
	RouteFunc:   AddRoute,
	RoutePrefix: RoutePrefix,
	// cobra flag
	FlagFunc: ServeFlag,
}

func ServeFlag(serveCmd *cobra.Command) {
	serveCmd.Flags().Bool(ID, DefaultEnable, Comment)
	serveCmd.Flags().String("ip2location.db", "DB11", "IP 数据库等级, 可选 DB1 DB3 DB5 DB9 DB11, 数字越大数据库内容越丰富, 相应数据库也就越大")
	serveCmd.Flags().String("ip2location.token", "", "ip2location lite token")
}
