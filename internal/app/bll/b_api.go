package bll

import (
	"context"

	"gin-admin-template/internal/app/schema"
)

// IApi 接口管理业务逻辑接口
type IApi interface {
	// 查询数据
	Query(ctx context.Context, params schema.ApiQueryParam, opts ...schema.ApiQueryOptions) (*schema.ApiQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, id string, opts ...schema.ApiGetOptions) (*schema.Api, error)
	// 创建数据
	Create(ctx context.Context, item schema.Api) (*schema.IDResult, error)
	// 更新数据
	Update(ctx context.Context, id string, item schema.Api) error
	// 删除数据
	Delete(ctx context.Context, id string) error
}
