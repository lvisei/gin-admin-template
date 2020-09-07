package api

import (
	"gin-admin-template/internal/app/bll"
	"gin-admin-template/internal/app/ginplus"
	"gin-admin-template/internal/app/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ResourceSet 注入Resource
var ResourceSet = wire.NewSet(wire.Struct(new(Resource), "*"))

// Resource 资源管理
type Resource struct {
	ResourceBll bll.IResource
}

// Query 查询数据
func (a *Resource) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ResourceQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.ResourceBll.Query(ctx, params)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResPage(c, result.Data, result.PageResult)
}

// QuerySelect 查询选择数据
func (a *Resource) QuerySelect(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ResourceQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}

	result, err := a.ResourceBll.Query(ctx, params, schema.ResourceQueryOptions{
		OrderFields: schema.NewOrderFields(schema.NewOrderField("`group`", schema.OrderByASC)),
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResList(c, result.Data)
}

// Get 查询指定数据
func (a *Resource) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ResourceBll.Get(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

// Create 创建数据
func (a *Resource) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Resource
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	item.Creator = ginplus.GetUserID(c)
	result, err := a.ResourceBll.Create(ctx, item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, result)
}

// Update 更新数据
func (a *Resource) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Resource
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	err := a.ResourceBll.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Delete 删除数据
func (a *Resource) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ResourceBll.Delete(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
