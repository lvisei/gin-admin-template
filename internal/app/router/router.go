package router

import (
	"gin-admin-template/internal/app/api"
	"gin-admin-template/pkg/auth"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var _ IRouter = (*Router)(nil)

// RouterSet 注入router
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

// Router 路由管理器
type Router struct {
	Auth           auth.Auther
	CasbinEnforcer *casbin.SyncedEnforcer
	DemoAPI        *api.Demo
	MockAPI        *api.Mock
	LoginAPI       *api.Login
	MenuAPI        *api.Menu
	RoleAPI        *api.Role
	UserAPI        *api.User
	SysAPI         *api.Sys
	ResourceAPI    *api.Resource
}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

// Prefixes 路由前缀列表
func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}
