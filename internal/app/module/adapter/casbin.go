package adapter

import (
	"context"
	"fmt"

	"gin-admin-template/internal/app/model/gorm/model"
	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/logger"
	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/google/wire"
)

var _ persist.Adapter = (*CasbinAdapter)(nil)

// CasbinAdapterSet 注入CasbinAdapter
var CasbinAdapterSet = wire.NewSet(wire.Struct(new(CasbinAdapter), "*"), wire.Bind(new(persist.Adapter), new(*CasbinAdapter)))

// CasbinAdapter casbin适配器
type CasbinAdapter struct {
	ResourceModel           *model.Resource
	RoleModel               *model.Role
	RoleMenuModel           *model.RoleMenu
	MenuActionResourceModel *model.MenuActionResource
	MenuResourceModel       *model.MenuResource
	UserModel               *model.User
	UserRoleModel           *model.UserRole
}

// LoadPolicy loads all policy rules from the storage.
func (a *CasbinAdapter) LoadPolicy(model casbinModel.Model) error {
	ctx := context.Background()
	err := a.loadRolePolicy(ctx, model)
	if err != nil {
		logger.WithContext(ctx).Errorf("Load casbin role policy error: %s", err.Error())
		return err
	}

	err = a.loadUserPolicy(ctx, model)
	if err != nil {
		logger.WithContext(ctx).Errorf("Load casbin user policy error: %s", err.Error())
		return err
	}

	return nil
}

// 加载角色策略(p,role_id,path,method)
func (a *CasbinAdapter) loadRolePolicy(ctx context.Context, m casbinModel.Model) error {
	resourceResult, err := a.ResourceModel.Query(ctx, schema.ResourceQueryParam{})
	if err != nil {
		return err
	} else if len(resourceResult.Data) == 0 {
		return nil
	}
	resources := resourceResult.Data.ToMap()

	roleResult, err := a.RoleModel.Query(ctx, schema.RoleQueryParam{
		Status: 1,
	})
	if err != nil {
		return err
	} else if len(roleResult.Data) == 0 {
		return nil
	}

	roleMenuResult, err := a.RoleMenuModel.Query(ctx, schema.RoleMenuQueryParam{})
	if err != nil {
		return err
	}
	mRoleMenus := roleMenuResult.Data.ToRoleIDMap()

	menuResourceResult, err := a.MenuResourceModel.Query(ctx, schema.MenuResourceQueryParam{})
	if err != nil {
		return err
	}
	mMenuResources := menuResourceResult.Data.ToMenuIDMap()

	menuActionResourceResult, err := a.MenuActionResourceModel.Query(ctx, schema.MenuActionResourceQueryParam{})
	if err != nil {
		return err
	}
	mActionMenuResources := menuActionResourceResult.Data.ToActionIDMap()

	getPolicyLine := func(roleId, resourceID string, mcache map[string]struct{}) string {
		if resource, ok := resources[resourceID]; ok {
			if resource.Path == "" || resource.Method == "" {
				return ""
			} else if _, ok := mcache[resource.Path+resource.Method]; ok {
				return ""
			}
			mcache[resource.Path+resource.Method] = struct{}{}
			line := fmt.Sprintf("p,%s,%s,%s", roleId, resource.Path, resource.Method)
			return line
		} else {
			return ""
		}
	}

	for _, item := range roleResult.Data {
		mcache := make(map[string]struct{})
		if roleMenus, ok := mRoleMenus[item.ID]; ok {
			for _, actionID := range roleMenus.ToActionIDs() {
				if menuActionResources, ok := mActionMenuResources[actionID]; ok {
					for _, menuActionResource := range menuActionResources {
						if line := getPolicyLine(item.ID, menuActionResource.ResourceID, mcache); line != "" {
							persist.LoadPolicyLine(line, m)
						}
					}
				}
			}
			for _, menuID := range roleMenus.ToMenuIDs() {
				if menuResources, ok := mMenuResources[menuID]; ok {
					for _, menuResource := range menuResources {
						if line := getPolicyLine(item.ID, menuResource.ResourceID, mcache); line != "" {
							persist.LoadPolicyLine(line, m)
						}
					}

				}
			}
		}
	}

	return nil
}

// 加载用户策略(g,user_id,role_id)
func (a *CasbinAdapter) loadUserPolicy(ctx context.Context, m casbinModel.Model) error {
	userResult, err := a.UserModel.Query(ctx, schema.UserQueryParam{
		Status: 1,
	})
	if err != nil {
		return err
	} else if len(userResult.Data) > 0 {
		userRoleResult, err := a.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{})
		if err != nil {
			return err
		}

		mUserRoles := userRoleResult.Data.ToUserIDMap()
		for _, uitem := range userResult.Data {
			if urs, ok := mUserRoles[uitem.ID]; ok {
				for _, ur := range urs {
					line := fmt.Sprintf("g,%s,%s", ur.UserID, ur.RoleID)
					persist.LoadPolicyLine(line, m)
				}
			}
		}
	}

	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *CasbinAdapter) SavePolicy(model casbinModel.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
