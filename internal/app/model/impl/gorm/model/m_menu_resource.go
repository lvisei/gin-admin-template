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

var _ model.IMenuResource = (*MenuResource)(nil)

// MenuResourceSet 注入MenuResource
var MenuResourceSet = wire.NewSet(wire.Struct(new(MenuResource), "*"), wire.Bind(new(model.IMenuResource), new(*MenuResource)))

// MenuResource 菜单资源存储
type MenuResource struct {
	DB *gorm.DB
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

	db := entity.GetMenuResourceDB(ctx, a.DB)
	if v := params.MenuID; v != "" {
		db= db.Where("menu_id=?", v)
	}
	if v := params.MenuIDs; len(v) > 0 {
		db		= entity.GetMenuActionDB(ctx, a.DB).Where("menu_id IN (?)", v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByASC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.MenuResources
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
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
	db := entity.GetMenuResourceDB(ctx, a.DB).Where("id=?", id)
	var item entity.MenuResource
	ok, err := FindOne(ctx, db, &item)
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
	result := entity.GetMenuResourceDB(ctx, a.DB).Create(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *MenuResource) Update(ctx context.Context, id string, item schema.MenuResource) error {
	eitem := entity.SchemaMenuResource(item).ToMenuResource()
	result := entity.GetMenuResourceDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *MenuResource) Delete(ctx context.Context, id string) error {
	result := entity.GetMenuResourceDB(ctx, a.DB).Where("id=?", id).Delete(entity.MenuResource{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteByMenuID 根据菜单ID删除数据
func (a *MenuResource) DeleteByMenuID(ctx context.Context, menuID string) error {
	result := entity.GetMenuResourceDB(ctx, a.DB).Where("menu_id=?", menuID).Delete(entity.MenuResource{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteByResourceID 根据菜单ID删除数据
func (a *MenuResource) DeleteByResourceID(ctx context.Context, resourceID string) error {
	result := entity.GetMenuResourceDB(ctx, a.DB).Where("resource_id=?", resourceID).Delete(entity.MenuResource{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
