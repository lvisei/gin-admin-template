package entity

import (
	"context"

	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/util/structure"
	"github.com/jinzhu/gorm"
)

// GetMenuDB 获取菜单存储
func GetMenuDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Menu))
}

// SchemaMenu 菜单对象
type SchemaMenu schema.Menu

// ToMenu 转换为菜单实体
func (a SchemaMenu) ToMenu() *Menu {
	item := new(Menu)
	structure.Copy(a, item)
	return item
}

// Menu 菜单实体
type Menu struct {
	Model
	Name       string  `gorm:"column:name;size:50;index;default:'';not null;"` // 菜单名称
	Sequence   int     `gorm:"column:sequence;index;default:0;not null;"`      // 排序值
	Icon       *string `gorm:"column:icon;size:255;"`                          // 菜单图标
	RouteName  *string `gorm:"column:route_name;size:255;"`                    // 路由名称
	RoutePath  *string `gorm:"column:route_path;size:255;"`                    // 路由地址
	Component  *string `gorm:"column:component;size:255;"`                     // 组件路径
	ParentID   *string `gorm:"column:parent_id;size:36;index;"`                // 父级内码
	ParentPath *string `gorm:"column:parent_path;size:518;index;"`             // 父级路径
	ShowStatus int     `gorm:"column:show_status;index;default:0;not null;"`   // 状态(1:显示 2:隐藏)
	Status     int     `gorm:"column:status;index;default:0;not null;"`        // 状态(1:启用 2:禁用)
	Memo       *string `gorm:"column:memo;size:1024;"`                         // 备注
	Creator    string  `gorm:"column:creator;size:36;"`                        // 创建人
}

// ToSchemaMenu 转换为菜单对象
func (a Menu) ToSchemaMenu() *schema.Menu {
	item := new(schema.Menu)
	structure.Copy(a, item)
	return item
}

// Menus 菜单实体列表
type Menus []*Menu

// ToSchemaMenus 转换为菜单对象列表
func (a Menus) ToSchemaMenus() []*schema.Menu {
	list := make([]*schema.Menu, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaMenu()
	}
	return list
}
