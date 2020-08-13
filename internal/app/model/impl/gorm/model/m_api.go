package model

import (
	"context"

	"gin-admin-template/internal/app/model"
	"gin-admin-template/internal/app/model/impl/gorm/entity"
	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/errors"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var _ model.IApi = (*Api)(nil)

// ApiSet 注入Api
var ApiSet = wire.NewSet(wire.Struct(new(Api), "*"), wire.Bind(new(model.IApi), new(*Api)))

// Api 接口管理存储
type Api struct {
	DB *gorm.DB
}

func (a *Api) getQueryOption(opts ...schema.ApiQueryOptions) schema.ApiQueryOptions {
	var opt schema.ApiQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Api) Query(ctx context.Context, params schema.ApiQueryParam, opts ...schema.ApiQueryOptions) (*schema.ApiQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetApiDB(ctx, a.DB)
	// TODO: 查询条件

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.Apis
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.ApiQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaApis(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Api) Get(ctx context.Context, id string, opts ...schema.ApiGetOptions) (*schema.Api, error) {
	db := entity.GetApiDB(ctx, a.DB).Where("id=?", id)
	var item entity.Api
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaApi(), nil
}

// Create 创建数据
func (a *Api) Create(ctx context.Context, item schema.Api) error {
	eitem := entity.SchemaApi(item).ToApi()
	result := entity.GetApiDB(ctx, a.DB).Create(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *Api) Update(ctx context.Context, id string, item schema.Api) error {
	eitem := entity.SchemaApi(item).ToApi()
	result := entity.GetApiDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *Api) Delete(ctx context.Context, id string) error {
	result := entity.GetApiDB(ctx, a.DB).Where("id=?", id).Delete(entity.Api{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
