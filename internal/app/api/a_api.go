package api

import (
	"gin-admin-template/internal/app/bll"
	"gin-admin-template/internal/app/ginplus"
	"gin-admin-template/internal/app/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ApiSet 注入Api
var ApiSet = wire.NewSet(wire.Struct(new(Api), "*"))

// Api 接口管理
type Api struct {
	ApiBll bll.IApi
}

// Query 查询数据
func (a *Api) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ApiQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.ApiBll.Query(ctx, params)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *Api) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ApiBll.Get(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

// Create 创建数据
func (a *Api) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Api
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	item.Creator = ginplus.GetUserID(c)
	result, err := a.ApiBll.Create(ctx, item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, result)
}

// Update 更新数据
func (a *Api) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Api
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	err := a.ApiBll.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Delete 删除数据
func (a *Api) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ApiBll.Delete(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
