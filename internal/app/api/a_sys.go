package api

import (
	"gin-admin-template/internal/app/ginplus"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// SysSet 注入Sys
var SysSet = wire.NewSet(wire.Struct(new(Sys), "*"))

// Demo 示例程序
type Sys struct{}

// LogCount 登录日志记录
func (a *Sys) LogCount(c *gin.Context) {

	data := map[string]interface{}{
		"online":    random(100, 500),
		"newVisits": random(160, 200),
		"totalUser": random(500, 1000),
		"messages":  random(10, 100),
	}

	ginplus.ResSuccess(c, data)
}
