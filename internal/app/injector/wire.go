// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"gin-admin-template/internal/app/api"
	// "gin-admin-template/internal/app/api/mock"
	"gin-admin-template/internal/app/bll/impl/bll"
	"gin-admin-template/internal/app/module/adapter"
	"gin-admin-template/internal/app/router"
	"github.com/google/wire"

	// mongoModel "gin-admin-template/internal/app/model/impl/mongo/model"
	gormModel "gin-admin-template/internal/app/model/impl/gorm/model"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	// 默认使用gorm存储注入，这里可使用 InitMongoDB & mongoModel.ModelSet 替换为 gorm 存储
	wire.Build(
		InitGormDB,
		gormModel.ModelSet,
		// InitMongoDB,
		// mongoModel.ModelSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		bll.BllSet,
		api.APISet,
		// mock.MockSet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
