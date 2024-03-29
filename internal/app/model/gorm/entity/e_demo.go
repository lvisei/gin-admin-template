package entity

import (
	"context"

	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/util/structure"
	"github.com/jinzhu/gorm"
)

// GetDemoDB 获取demo存储
func GetDemoDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Demo))
}

// SchemaDemo demo对象
type SchemaDemo schema.Demo

// ToDemo 转换为demo实体
func (a SchemaDemo) ToDemo() *Demo {
	item := new(Demo)
	structure.Copy(a, item)
	return item
}

// Demo demo实体
type Demo struct {
	Model
	Code    string  `gorm:"column:code;size:50;index;default:'';not null;"`  // 编号
	Name    string  `gorm:"column:name;size:100;index;default:'';not null;"` // 名称
	Memo    *string `gorm:"column:memo;size:200;"`                           // 备注
	Status  int     `gorm:"column:status;index;default:0;not null;"`         // 状态(1:启用 2:停用)
	Creator string  `gorm:"column:creator;size:36;"`                         // 创建者
}

// ToSchemaDemo 转换为demo对象
func (a Demo) ToSchemaDemo() *schema.Demo {
	item := new(schema.Demo)
	structure.Copy(a, item)
	return item
}

// Demos demo列表
type Demos []*Demo

// ToSchemaDemos 转换为demo对象列表
func (a Demos) ToSchemaDemos() []*schema.Demo {
	list := make([]*schema.Demo, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaDemo()
	}
	return list
}
