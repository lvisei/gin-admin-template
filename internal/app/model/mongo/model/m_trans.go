package model

import (
	"context"
	"gin-admin-template/internal/app/contextx"

	"gin-admin-template/pkg/errors"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// TransSet 注入Trans
var TransSet = wire.NewSet(wire.Struct(new(Trans), "*"))

// Trans 事务管理
type Trans struct {
	Client *mongo.Client
}

// Exec 执行事务
func (a *Trans) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := contextx.FromTrans(ctx); ok {
		return fn(ctx)
	}

	session, err := a.Client.StartSession()
	if err != nil {
		return errors.WithStack(err)
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		err := fn(contextx.NewTrans(sessCtx, true))
		return nil, err
	})

	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
