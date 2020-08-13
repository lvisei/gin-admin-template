package schema

import (
	"time"

	"gin-admin-template/pkg/util"
)

// Api 接口管理对象
type Api struct {
	ID          string    `json:"id"`                        // 唯一标识
	Group       string    `json:"group" binding:"required"`  // 接口组
	Path        string    `json:"path" binding:"required"`   // 资源请求路径（支持/:id匹配）
	Method      string    `json:"method" binding:"required"` // 资源请求方式(支持正则)
	Description string    `json:"description"`               // 接口描述
	Creator     string    `json:"creator"`                   // 创建者
	CreatedAt   time.Time `json:"created_at"`                // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`                // 更新时间

}

func (a *Api) String() string {
	return util.JSONMarshalToString(a)
}

// ApiQueryParam 查询条件
type ApiQueryParam struct {
	PaginationParam
}

// ApiQueryOptions 查询可选参数项
type ApiQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// ApiGetOptions Get查询可选参数项
type ApiGetOptions struct {
}

// ApiQueryResult 查询结果
type ApiQueryResult struct {
	Data       Apis
	PageResult *PaginationResult
}

// Apis 接口管理列表
type Apis []*Api
