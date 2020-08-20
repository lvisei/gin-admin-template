package model

import (
	"context"
	"time"

	"gin-admin-template/internal/app/model"
	"gin-admin-template/internal/app/model/impl/mongo/entity"
	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/errors"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ model.IMenuResource = (*MenuResource)(nil)

// MenuResourceSet 注入MenuResource
var MenuResourceSet = wire.NewSet(wire.Struct(new(MenuResource), "*"), wire.Bind(new(model.IMenuResource), new(*MenuResource)))

// MenuResource 菜单资源存储
type MenuResource struct {
	Client *mongo.Client
}

func (a *MenuResource) getQueryOption(opts ...schema.MenuResourceQueryOptions) schema.MenuResourceQueryOptions {
	var opt schema.MenuResourceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *MenuResource) Query(ctx context.Context, params schema.MenuResourceQueryParam, opts ...schema.MenuResourceQueryOptions) (*schema.MenuResourceQueryResult, error) {
	opt := a.getQueryOption(opts...)

	c := entity.GetMenuResourceCollection(ctx, a.Client)
	filter := DefaultFilter(ctx)
	menuIDs := params.MenuIDs
	if v := params.MenuID; v != "" {
		menuIDs = append(menuIDs, v)
	}
	if v := menuIDs; len(v) > 0 {
		filter = append(filter, Filter("menu_id", bson.M{"$in": menuIDs}))
	}
	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("_id", schema.OrderByASC))

	var list entity.MenuResources
	pr, err := WrapPageQuery(ctx, c, params.PaginationParam, filter, &list, options.Find().SetSort(ParseOrder(opt.OrderFields)))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.MenuResourceQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaMenuResources(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *MenuResource) Get(ctx context.Context, id string, opts ...schema.MenuResourceGetOptions) (*schema.MenuResource, error) {
	c := entity.GetMenuResourceCollection(ctx, a.Client)
	filter := DefaultFilter(ctx, Filter("_id", id))
	var item entity.MenuResource
	ok, err := FindOne(ctx, c, filter, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaMenuResource(), nil
}

// Create 创建数据
func (a *MenuResource) Create(ctx context.Context, item schema.MenuResource) error {
	eitem := entity.SchemaMenuResource(item).ToMenuResource()
	eitem.CreatedAt = time.Now()
	eitem.UpdatedAt = time.Now()
	c := entity.GetMenuResourceCollection(ctx, a.Client)
	err := Insert(ctx, c, eitem)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *MenuResource) Update(ctx context.Context, id string, item schema.MenuResource) error {
	eitem := entity.SchemaMenuResource(item).ToMenuResource()
	eitem.UpdatedAt = time.Now()
	c := entity.GetMenuResourceCollection(ctx, a.Client)
	err := Update(ctx, c, DefaultFilter(ctx, Filter("_id", id)), eitem)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *MenuResource) Delete(ctx context.Context, id string) error {
	c := entity.GetMenuResourceCollection(ctx, a.Client)
	err := Delete(ctx, c, DefaultFilter(ctx, Filter("_id", id)))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteByMenuID 根据菜单ID删除数据
func (a *MenuResource) DeleteByMenuID(ctx context.Context, menuID string) error {
	c := entity.GetMenuResourceCollection(ctx, a.Client)
	err := DeleteMany(ctx, c, DefaultFilter(ctx, Filter("menu_id", menuID)))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteByResourceID 根据菜单ID删除数据
func (a *MenuResource) DeleteByResourceID(ctx context.Context, resourceID string) error {
	c := entity.GetMenuResourceCollection(ctx, a.Client)
	err := DeleteMany(ctx, c, DefaultFilter(ctx, Filter("resource_id", resourceID)))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
