package bll

import (
	"context"

	"gin-admin-template/internal/app/bll"
	"gin-admin-template/internal/app/iutil"
	"gin-admin-template/internal/app/model"
	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/errors"
	"github.com/google/wire"
)

var _ bll.IResource = (*Resource)(nil)

// ResourceSet 注入Resource
var ResourceSet = wire.NewSet(wire.Struct(new(Resource), "*"), wire.Bind(new(bll.IResource), new(*Resource)))

// Resource 资源管理
type Resource struct {
	ResourceModel model.IResource
}

// Query 查询数据
func (a *Resource) Query(ctx context.Context, params schema.ResourceQueryParam, opts ...schema.ResourceQueryOptions) (*schema.ResourceQueryResult, error) {
	return a.ResourceModel.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *Resource) Get(ctx context.Context, id string, opts ...schema.ResourceGetOptions) (*schema.Resource, error) {
	item, err := a.ResourceModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *Resource) Create(ctx context.Context, item schema.Resource) (*schema.IDResult, error) {
	item.ID = iutil.NewID()
	err := a.ResourceModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

// Update 更新数据
func (a *Resource) Update(ctx context.Context, id string, item schema.Resource) error {
	oldItem, err := a.ResourceModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt

	return a.ResourceModel.Update(ctx, id, item)
}

// Delete 删除数据
func (a *Resource) Delete(ctx context.Context, id string) error {
	oldItem, err := a.ResourceModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.ResourceModel.Delete(ctx, id)
}
