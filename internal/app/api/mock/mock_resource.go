package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ResourceSet 注入Resource
var ResourceSet = wire.NewSet(wire.Struct(new(Resource), "*"))

// Resource 资源管理
type Resource struct {
}

// Query 查询数据
// @Tags 资源管理
// @Summary 查询数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param queryValue query string false "查询值"
// @Param group query string false "接口组"
// @Param path query string false "请求路径"
// @Param method query string false "请求方式"
// @Success 200 {array} schema.Resource "查询结果：{list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/resources [get]
func (a *Resource) Query(c *gin.Context) {
}

// QuerySelect 查询选择数据
// @Tags 资源管理
// @Summary 查询选择数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param queryValue query string false "查询值"
// @Success 200 {array} schema.Resource "查询结果：{list:资源列表}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/resources.select [get]
func (a *Resource) QuerySelect(c *gin.Context) {
}

// Get 查询指定数据
// @Tags 资源管理
// @Summary 查询指定数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.Resource
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.ErrorResult "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/resources/{id} [get]
func (a *Resource) Get(c *gin.Context) {
}

// Create 创建数据
// @Tags 资源管理
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.ResourceCreateParams true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/resources [post]
func (a *Resource) Create(c *gin.Context) {
}

// Update 更新数据
// @Tags 资源管理
// @Summary 更新数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Param body body schema.ResourceCreateParams true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/resources/{id} [put]
func (a *Resource) Update(c *gin.Context) {
}

// Delete 删除数据
// @Tags 资源管理
// @Summary 删除数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/resources/{id} [delete]
func (a *Resource) Delete(c *gin.Context) {
}
