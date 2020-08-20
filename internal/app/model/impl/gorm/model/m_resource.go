package model

import (
	"context"

	"gin-admin-template/internal/app/model"
	"gin-admin-template/internal/app/model/impl/gorm/entity"
	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/errors"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var _ model.IResource = (*Resource)(nil)

// ResourceSet 注入Resource
var ResourceSet = wire.NewSet(wire.Struct(new(Resource), "*"), wire.Bind(new(model.IResource), new(*Resource)))

// Resource 资源管理存储
type Resource struct {
	DB *gorm.DB
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

	db := entity.GetResourceDB(ctx, a.DB)
	if v := params.Group; v != "" {
		db = db.Where("group=?", v)
	}
	if v := params.Path; v != "" {
		db = db.Where("path=?", v)
	}
	if v := params.Method; v != "" {
		db = db.Where("method=?", v)
	}
	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("group LIKE ? OR description LIKE ?", v, v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.Resources
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
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
	db := entity.GetResourceDB(ctx, a.DB).Where("id=?", id)
	var item entity.Resource
	ok, err := FindOne(ctx, db, &item)
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
	result := entity.GetResourceDB(ctx, a.DB).Create(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *Resource) Update(ctx context.Context, id string, item schema.Resource) error {
	eitem := entity.SchemaResource(item).ToResource()
	result := entity.GetResourceDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *Resource) Delete(ctx context.Context, id string) error {
	result := entity.GetResourceDB(ctx, a.DB).Where("id=?", id).Delete(entity.Resource{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
