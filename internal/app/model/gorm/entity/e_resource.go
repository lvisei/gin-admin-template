package entity

import (
	"context"

	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/util/structure"
	"github.com/jinzhu/gorm"
)

// GetResourceDB 获取Resource存储
func GetResourceDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Resource))
}

// SchemaResource 资源管理对象
type SchemaResource schema.Resource

// ToResource 转换为实体
func (a SchemaResource) ToResource() *Resource {
	item := new(Resource)
	structure.Copy(a, item)
	return item
}

// Resource 资源管理实体
type Resource struct {
	Model
	Group       string `gorm:"column:group;size:50;index;default:'';not null;"` // 资源组
	Path        string `gorm:"column:path;size:100;"`                           // 资源请求路径（支持/:id匹配）
	Method      string `gorm:"column:method;size:100;"`                         // 资源请求方式(支持正则)
	Description string `gorm:"column:description;size:1024;"`                   // 资源描述
	Creator     string `gorm:"column:creator;"`                                 // 创建者
}

// ToSchemaResource 转换为demo对象
func (a Resource) ToSchemaResource() *schema.Resource {
	item := new(schema.Resource)
	structure.Copy(a, item)
	return item
}

// Resources 资源管理实体列表
type Resources []*Resource

// ToSchemaResources 转换为对象列表
func (a Resources) ToSchemaResources() schema.Resources {
	list := make(schema.Resources, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaResource()
	}
	return list
}
