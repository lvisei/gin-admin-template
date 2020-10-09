package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"gin-admin-template/internal/app/model/mongo/entity"
	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/errors"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ResourceSet 注入Resource
var ResourceSet = wire.NewSet(wire.Struct(new(Resource), "*"))

// Resource 资源管理存储
type Resource struct {
	Client *mongo.Client
}

func (a *Resource) getQueryOption(opts ...schema.ResourceQueryOptions) schema.ResourceQueryOptions {
	var opt schema.ResourceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Resource) Query(ctx context.Context, params schema.ResourceQueryParam, opts ...schema.ResourceQueryOptions) (*schema.ResourceQueryResult, error) {
	opt := a.getQueryOption(opts...)

	c := entity.GetResourceCollection(ctx, a.Client)
	filter := DefaultFilter(ctx)
	if v := params.Group; v != "" {
		filter = append(filter, Filter("group", v))
	}
	if v := params.Path; v != "" {
		filter = append(filter, Filter("path", v))
	}
	if v := params.Method; v != "" {
		filter = append(filter, Filter("method", v))
	}
	if v := params.QueryValue; v != "" {
		filter = append(filter, Filter("$or", bson.A{
			OrRegexFilter("group", v),
			OrRegexFilter("description", v),
		}))
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("_id", schema.OrderByDESC))

	var list entity.Resources
	pr, err := WrapPageQuery(ctx, c, params.PaginationParam, filter, &list, options.Find().SetSort(ParseOrder(opt.OrderFields)))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.ResourceQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaResources(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Resource) Get(ctx context.Context, id string, opts ...schema.ResourceGetOptions) (*schema.Resource, error) {
	c := entity.GetResourceCollection(ctx, a.Client)
	filter := DefaultFilter(ctx, Filter("_id", id))
	var item entity.Resource
	ok, err := FindOne(ctx, c, filter, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaResource(), nil
}

// Create 创建数据
func (a *Resource) Create(ctx context.Context, item schema.Resource) error {
	eitem := entity.SchemaResource(item).ToResource()
	eitem.CreatedAt = time.Now()
	eitem.UpdatedAt = time.Now()
	c := entity.GetResourceCollection(ctx, a.Client)
	err := Insert(ctx, c, eitem)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *Resource) Update(ctx context.Context, id string, item schema.Resource) error {
	eitem := entity.SchemaResource(item).ToResource()
	eitem.UpdatedAt = time.Now()
	c := entity.GetResourceCollection(ctx, a.Client)
	err := Update(ctx, c, DefaultFilter(ctx, Filter("_id", id)), eitem)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *Resource) Delete(ctx context.Context, id string) error {
	c := entity.GetResourceCollection(ctx, a.Client)
	err := Delete(ctx, c, DefaultFilter(ctx, Filter("_id", id)))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
