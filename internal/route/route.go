package route

import (
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func Register(group string, cb func(gin.IRouter)) {
	subRoute := server.Group(group)
	cb(subRoute)
}

func init() {
	server = gin.Default()
}

func Run(addr ...string) {
	server.Run(addr...)
}
