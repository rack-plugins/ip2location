package ip2location

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("allservices") {
		return
	}
	g.GET("/help/ip2location", help)

	r := g.Group(RoutePrefix)
	r.GET("/:ip", IpQuery)
	r.POST("/:ip", IpQuery)
}
