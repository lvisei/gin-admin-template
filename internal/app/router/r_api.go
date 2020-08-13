package router

import (
	"gin-admin-template/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	g.Use(middleware.UserAuthMiddleware(a.Auth,
		middleware.AllowPathPrefixSkipper("/api/v1/pub/login"),
	))

	g.Use(middleware.CasbinMiddleware(a.CasbinEnforcer,
		middleware.AllowPathPrefixSkipper("/api/v1/pub"),
	))

	g.Use(middleware.RateLimiterMiddleware())

	v1 := g.Group("/v1")
	{
		pub := v1.Group("/pub")
		{
			gLogin := pub.Group("login")
			{
				gLogin.GET("captchaid", a.LoginAPI.GetCaptcha)
				gLogin.GET("captcha", a.LoginAPI.ResCaptcha)
				gLogin.POST("", a.LoginAPI.Login)
				gLogin.POST("exit", a.LoginAPI.Logout)
			}

			gCurrent := pub.Group("current")
			{
				gCurrent.PUT("password", a.LoginAPI.UpdatePassword)
				gCurrent.PUT("user", a.LoginAPI.UpdateUserInfo)
				gCurrent.GET("user", a.LoginAPI.GetUserInfo)
				gCurrent.GET("menutree", a.LoginAPI.QueryUserMenuTree)
			}
			pub.POST("/refresh-token", a.LoginAPI.RefreshToken)
		}

		gMock := v1.Group("mock")
		{
			gMock.GET("users", a.MockAPI.UserQuery)
		}

		gDemo := v1.Group("demos")
		{
			gDemo.GET("", a.DemoAPI.Query)
			gDemo.GET(":id", a.DemoAPI.Get)
			gDemo.POST("", a.DemoAPI.Create)
			gDemo.PUT(":id", a.DemoAPI.Update)
			gDemo.DELETE(":id", a.DemoAPI.Delete)
			gDemo.PATCH(":id/enable", a.DemoAPI.Enable)
			gDemo.PATCH(":id/disable", a.DemoAPI.Disable)
		}

		gMenu := v1.Group("menus")
		{
			gMenu.GET("", a.MenuAPI.Query)
			gMenu.GET(":id", a.MenuAPI.Get)
			gMenu.POST("", a.MenuAPI.Create)
			gMenu.PUT(":id", a.MenuAPI.Update)
			gMenu.DELETE(":id", a.MenuAPI.Delete)
			gMenu.PATCH(":id/enable", a.MenuAPI.Enable)
			gMenu.PATCH(":id/disable", a.MenuAPI.Disable)
		}
		v1.GET("/menus.tree", a.MenuAPI.QueryTree)

		gRole := v1.Group("roles")
		{
			gRole.GET("", a.RoleAPI.Query)
			gRole.GET(":id", a.RoleAPI.Get)
			gRole.POST("", a.RoleAPI.Create)
			gRole.PUT(":id", a.RoleAPI.Update)
			gRole.DELETE(":id", a.RoleAPI.Delete)
			gRole.PATCH(":id/enable", a.RoleAPI.Enable)
			gRole.PATCH(":id/disable", a.RoleAPI.Disable)
		}
		v1.GET("/roles.select", a.RoleAPI.QuerySelect)

		gUser := v1.Group("users")
		{
			gUser.GET("", a.UserAPI.Query)
			gUser.GET(":id", a.UserAPI.Get)
			gUser.POST("", a.UserAPI.Create)
			gUser.PUT(":id", a.UserAPI.Update)
			gUser.PUT(":id/rest", a.UserAPI.ResetPassword)
			gUser.DELETE(":id", a.UserAPI.Delete)
			gUser.PATCH(":id/enable", a.UserAPI.Enable)
			gUser.PATCH(":id/disable", a.UserAPI.Disable)
		}

		gSys := v1.Group("sys")
		{
			gSys.GET("logs/count", a.SysAPI.LogCount)
		}

		gApi := v1.Group("apis")
		{
			gApi.GET("", a.ApiAPI.Query)
			gApi.GET(":id", a.ApiAPI.Get)
			gApi.POST("", a.ApiAPI.Create)
			gApi.PUT(":id", a.ApiAPI.Update)
			gApi.DELETE(":id", a.ApiAPI.Delete)
		}

	}
}
