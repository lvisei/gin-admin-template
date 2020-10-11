package bll

import (
	"context"
	"github.com/casbin/casbin/v2"

	"gin-admin-template/internal/app/model/gorm/model"
	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/errors"
	"gin-admin-template/pkg/util/uuid"
	"github.com/google/wire"
)

// ResourceSet 注入Resource
var ResourceSet = wire.NewSet(wire.Struct(new(Resource), "*"))

// Resource 资源管理
type Resource struct {
	Enforcer      *casbin.SyncedEnforcer
	ResourceModel *model.Resource
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

func (a *Resource) checkResource(ctx context.Context, item schema.Resource) error {
	result, err := a.ResourceModel.Query(ctx, schema.ResourceQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		Path:   item.Path,
		Method: item.Method,
		Group:  item.Group,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("资源已经存在")
	}
	return nil
}

// Create 创建数据
func (a *Resource) Create(ctx context.Context, item schema.Resource) (*schema.IDResult, error) {
	if err := a.checkResource(ctx, item); err != nil {
		return nil, err
	}

	item.ID = uuid.MustString()
	err := a.ResourceModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	LoadCasbinPolicy(ctx, a.Enforcer)
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

	err = a.ResourceModel.Update(ctx, id, item)
	if err != nil {
		return err
	}

	LoadCasbinPolicy(ctx, a.Enforcer)
	return nil
}

// Delete 删除数据
func (a *Resource) Delete(ctx context.Context, id string) error {
	oldItem, err := a.ResourceModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	err = a.ResourceModel.Delete(ctx, id)
	if err != nil {
		return err
	}

	LoadCasbinPolicy(ctx, a.Enforcer)
	return nil
}
