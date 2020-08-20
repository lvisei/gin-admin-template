package entity

import (
	"context"

	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetMenuResourceCollection 获取MenuResource集合
func GetMenuResourceCollection(ctx context.Context, cli *mongo.Client) *mongo.Collection {
	return GetCollection(ctx, cli, MenuResource{})
}

// SchemaMenuResource 菜单资源
type SchemaMenuResource schema.MenuResource

// ToMenuResource 转换为实体
func (a SchemaMenuResource) ToMenuResource() *MenuResource {
	item := new(MenuResource)
	util.StructMapToStruct(a, item)
	return item
}

// MenuResource 菜单资源实体
type MenuResource struct {
	Model      `bson:",inline"`
	MenuID     string `bson:"menu_id"`     // 菜单ID
	ResourceID string `bson:"resource_id"` // 资源ID
}

// CollectionName 集合名
func (a MenuResource) CollectionName() string {
	return a.Model.CollectionName("menu_resource")
}

// CreateIndexes 创建索引
func (a MenuResource) CreateIndexes(ctx context.Context, cli *mongo.Client) error {
	return a.Model.CreateIndexes(ctx, cli, a, []mongo.IndexModel{
		{Keys: bson.M{"creator": 1}},
	})
}

// ToSchemaMenuResource 转换为对象
func (a MenuResource) ToSchemaMenuResource() *schema.MenuResource {
	item := new(schema.MenuResource)
	util.StructMapToStruct(a, item)
	return item
}

// MenuResources 菜单资源实体列表
type MenuResources []*MenuResource

// ToSchemaMenuResources 转换为菜单资源对象列表
func (a MenuResources) ToSchemaMenuResources() schema.MenuResources {
	list := make(schema.MenuResources, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaMenuResource()
	}
	return list
}
