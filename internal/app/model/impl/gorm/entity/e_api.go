package entity

import (
	"context"

	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/util"
	"github.com/jinzhu/gorm"
)

// GetApiDB 获取Api存储
func GetApiDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Api))
}

// SchemaApi 接口管理对象
type SchemaApi schema.Api

// ToApi 转换为实体
func (a SchemaApi) ToApi() *Api {
	item := new(Api)
	util.StructMapToStruct(a, item)
	return item
}

// Api 接口管理实体
type Api struct {
	Model
	Group       string `gorm:"column:group;size:50;index;"`   // 接口组
	Path        string `gorm:"column:path;size:100;"`         // 资源请求路径（支持/:id匹配）
	Method      string `gorm:"column:method;size:100;"`       // 资源请求方式(支持正则)
	Description string `gorm:"column:description;size:1024;"` // 接口描述
	Creator     string `gorm:"column:creator;"`               // 创建者

}

// ToSchemaApi 转换为demo对象
func (a Api) ToSchemaApi() *schema.Api {
	item := new(schema.Api)
	util.StructMapToStruct(a, item)
	return item
}

// Apis 接口管理实体列表
type Apis []*Api

// ToSchemaApis 转换为对象列表
func (a Apis) ToSchemaApis() schema.Apis {
	list := make(schema.Apis, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaApi()
	}
	return list
}
