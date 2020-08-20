package entity

import (
	"context"

	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/util"
	"github.com/jinzhu/gorm"
)

// GetMenuResourceDB 获取MenuResource存储
func GetMenuResourceDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(MenuResource))
}

// SchemaMenuResource 菜单资源对象
type SchemaMenuResource schema.MenuResource

// ToMenuResource 转换为实体
func (a SchemaMenuResource) ToMenuResource() *MenuResource {
	item := new(MenuResource)
	util.StructMapToStruct(a, item)
	return item
}

// MenuResource 菜单资源实体
type MenuResource struct {
	Model
	MenuID     string `gorm:"column:menu_id;size:36;index;default:'';not null;"`     // 菜单ID
	ResourceID string `gorm:"column:resource_id;size:36;"` // 资源ID
}

// ToSchemaMenuResource 转换为demo对象
func (a MenuResource) ToSchemaMenuResource() *schema.MenuResource {
	item := new(schema.MenuResource)
	util.StructMapToStruct(a, item)
	return item
}

// MenuResources 菜单资源实体列表
type MenuResources []*MenuResource

// ToSchemaMenuResources 转换为对象列表
func (a MenuResources) ToSchemaMenuResources() schema.MenuResources {
	list := make(schema.MenuResources, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaMenuResource()
	}
	return list
}
