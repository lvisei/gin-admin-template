package schema

import (
	"time"

	"gin-admin-template/pkg/util"
)

// ResourceCreateParams 新增参数
type ResourceCreateParams struct {
	Group       string `json:"group" binding:"required"`  // 资源组
	Path        string `json:"path" binding:"required"`   // 资源请求路径（支持/:id匹配）
	Method      string `json:"method" binding:"required"` // 资源请求方式(支持正则)
	Description string `json:"description"`               // 资源描述
}

// Resource 资源管理对象
type Resource struct {
	ResourceCreateParams
	ID          string    `json:"id"`                        // 唯一标识
	Creator     string    `json:"creator"`                   // 创建者
	CreatedAt   time.Time `json:"created_at"`                // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`                // 更新时间
}

func (a *Resource) String() string {
	return util.JSONMarshalToString(a)
}

// ResourceQueryParam 查询条件
type ResourceQueryParam struct {
	PaginationParam
	QueryValue string `form:"queryValue"` // 模糊查询
	Group      string `form:"group"`      // 接口组
	Path       string `form:"path"`       // 资源请求路径（支持/:id匹配）
	Method     string `form:"method"`     // 资源请求方式(支持正则)
}

// ResourceQueryOptions 查询可选参数项
type ResourceQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// ResourceGetOptions Get查询可选参数项
type ResourceGetOptions struct {
}

// ResourceQueryResult 查询结果
type ResourceQueryResult struct {
	Data       Resources
	PageResult *PaginationResult
}

// Resources 资源管理列表
type Resources []*Resource
