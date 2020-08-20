package model

import (
	"context"

	"gin-admin-template/internal/app/schema"
)

// IMenuResource 菜单资源存储接口
type IMenuResource interface {
	// 查询数据
	Query(ctx context.Context, params schema.MenuResourceQueryParam, opts ...schema.MenuResourceQueryOptions) (*schema.MenuResourceQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, id string, opts ...schema.MenuResourceGetOptions) (*schema.MenuResource, error)
	// 创建数据
	Create(ctx context.Context, item schema.MenuResource) error
	// 更新数据
	Update(ctx context.Context, id string, item schema.MenuResource) error
	// 删除数据
	Delete(ctx context.Context, id string) error
	// 根据菜单ID删除数据
	DeleteByMenuID(ctx context.Context, menuID string) error
	// 根据资源ID删除数据
	DeleteByResourceID(ctx context.Context, resourceID string) error
}
