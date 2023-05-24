package ip2location

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func help(ctx *gin.Context) {
	ctx.String(http.StatusOK, `GET|POST /ip2location/ip/x.x.x.x
`)
}
