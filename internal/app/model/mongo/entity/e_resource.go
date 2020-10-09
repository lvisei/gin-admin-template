package entity

import (
	"context"

	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/util/structure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetResourceCollection 获取Resource集合
func GetResourceCollection(ctx context.Context, cli *mongo.Client) *mongo.Collection {
	return GetCollection(ctx, cli, Resource{})
}

// SchemaResource 资源管理
type SchemaResource schema.Resource

// ToResource 转换为实体
func (a SchemaResource) ToResource() *Resource {
	item := new(Resource)
	structure.Copy(a, item)
	return item
}

// Resource 资源管理实体
type Resource struct {
	Model       `bson:",inline"`
	Group       string `bson:"group"`       // 资源组
	Path        string `bson:"path"`        // 资源请求路径（支持/:id匹配）
	Method      string `bson:"method"`      // 资源请求方式(支持正则)
	Description string `bson:"description"` // 接口描述
	Creator     string `bson:"creator"`     // 创建者

}

// CollectionName 集合名
func (a Resource) CollectionName() string {
	return a.Model.CollectionName("resource")
}

// CreateIndexes 创建索引
func (a Resource) CreateIndexes(ctx context.Context, cli *mongo.Client) error {
	return a.Model.CreateIndexes(ctx, cli, a, []mongo.IndexModel{
		{Keys: bson.M{"creator": 1}},
	})
}

// ToSchemaResource 转换为对象
func (a Resource) ToSchemaResource() *schema.Resource {
	item := new(schema.Resource)
	structure.Copy(a, item)
	return item
}

// Resources 资源管理实体列表
type Resources []*Resource

// ToSchemaResources 转换为资源管理对象列表
func (a Resources) ToSchemaResources() schema.Resources {
	list := make(schema.Resources, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaResource()
	}
	return list
}
