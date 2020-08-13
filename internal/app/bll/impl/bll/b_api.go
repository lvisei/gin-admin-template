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

var _ bll.IApi = (*Api)(nil)

// ApiSet 注入Api
var ApiSet = wire.NewSet(wire.Struct(new(Api), "*"), wire.Bind(new(bll.IApi), new(*Api)))

// Api 接口管理
type Api struct {
	ApiModel model.IApi
}

// Query 查询数据
func (a *Api) Query(ctx context.Context, params schema.ApiQueryParam, opts ...schema.ApiQueryOptions) (*schema.ApiQueryResult, error) {
	return a.ApiModel.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *Api) Get(ctx context.Context, id string, opts ...schema.ApiGetOptions) (*schema.Api, error) {
	item, err := a.ApiModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *Api) Create(ctx context.Context, item schema.Api) (*schema.IDResult, error) {
	item.ID = iutil.NewID()
	err := a.ApiModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

// Update 更新数据
func (a *Api) Update(ctx context.Context, id string, item schema.Api) error {
	oldItem, err := a.ApiModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt

	return a.ApiModel.Update(ctx, id, item)
}

// Delete 删除数据
func (a *Api) Delete(ctx context.Context, id string) error {
	oldItem, err := a.ApiModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.ApiModel.Delete(ctx, id)
}
