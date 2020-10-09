// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"gin-admin-template/internal/app/api"
	bll2 "gin-admin-template/internal/app/bll"

	"gin-admin-template/internal/app/module/adapter"
	"gin-admin-template/internal/app/router"
	"github.com/google/wire"

	// mongoModel "gin-admin-template/internal/app/model/impl/mongo/model"
	gormModel "gin-admin-template/internal/app/model/gorm/model"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	// 默认使用gorm存储注入，这里可使用 InitMongoDB & mongoModel.ModelSet 替换为 gorm 存储
	wire.Build(
		// mock.MockSet,
		InitGormDB,
		gormModel.ModelSet,
		// InitMongoDB,
		// mongoModel.ModelSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		bll2.BllSet,
		api.APISet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
