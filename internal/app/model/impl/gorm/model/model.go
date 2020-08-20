package model

import "github.com/google/wire"

// ModelSet model注入
var ModelSet = wire.NewSet(
	DemoSet,
	MenuActionResourceSet,
	MenuResourceSet,
	MenuActionSet,
	MenuSet,
	RoleMenuSet,
	RoleSet,
	TransSet,
	UserRoleSet,
	UserSet,
	ApiSet,
	ResourceSet,
)
