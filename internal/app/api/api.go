package api

import "github.com/google/wire"

// APISet 注入api
var APISet = wire.NewSet(
	MockSet,
	DemoSet,
	LoginSet,
	MenuSet,
	RoleSet,
	UserSet,
	SysSet,
	ApiSet,
)
