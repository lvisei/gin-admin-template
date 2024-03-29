package schema

import (
	"gin-admin-template/pkg/util/json"
	"strings"
	"time"
)

// Menu 菜单对象
type Menu struct {
	ID         string        `json:"id"`                                        // 唯一标识
	Name       string        `json:"name" binding:"required"`                   // 菜单名称
	Sequence   int           `json:"sequence"`                                  // 排序值
	Icon       string        `json:"icon"`                                      // 菜单图标
	RouteName  string        `json:"routeName"`                                 // 路由名称
	RoutePath  string        `json:"routePath"`                                 // 路由地址
	Component  string        `json:"component"`                                 // 组件路径
	ParentID   string        `json:"parentId"`                                  // 父级ID
	ParentPath string        `json:"parentPath"`                                // 父级路径
	ShowStatus int           `json:"showStatus" binding:"required,max=2,min=1"` // 显示状态(1:显示 2:隐藏)
	Status     int           `json:"status" binding:"required,max=2,min=1"`     // 状态(1:启用 2:禁用)
	Memo       string        `json:"memo"`                                      // 备注
	Creator    string        `json:"creator"`                                   // 创建者
	CreatedAt  time.Time     `json:"createdAt"`                                 // 创建时间
	UpdatedAt  time.Time     `json:"updatedAt"`                                 // 更新时间
	Actions    MenuActions   `json:"actions"`                                   // 动作列表
	Resources  MenuResources `json:"resources"`                                 // 资源列表
}

func (a *Menu) String() string {
	return json.MarshalToString(a)
}

// MenuCreateParams 新增参数
type MenuCreateParams struct {
	Name       string        `json:"name" binding:"required"`                   // 菜单名称
	Sequence   int           `json:"sequence"`                                  // 排序值
	Icon       string        `json:"icon"`                                      // 菜单图标
	RouteName  string        `json:"routeName"`                                 // 路由名称
	RoutePath  string        `json:"routePath"`                                 // 路由地址
	Component  string        `json:"component"`                                 // 组件路径
	ParentID   string        `json:"parentId"`                                  // 父级ID
	ShowStatus int           `json:"showStatus" binding:"required,max=2,min=1"` // 显示状态(1:显示 2:隐藏)
	Status     int           `json:"status" binding:"required,max=2,min=1"`     // 状态(1:启用 2:禁用)
	Memo       string        `json:"memo"`                                      // 备注
	Actions    MenuActions   `json:"actions"`                                   // 动作列表
	Resources  MenuResources `json:"resources"`                                 // 资源列表
}

// MenuQueryParam 查询条件
type MenuQueryParam struct {
	PaginationParam
	IDs              []string `form:"-"`          // 唯一标识列表
	Name             string   `form:"-"`          // 菜单名称
	PrefixParentPath string   `form:"-"`          // 父级路径(前缀模糊查询)
	QueryValue       string   `form:"queryValue"` // 模糊查询
	ParentID         *string  `form:"parentID"`   // 父级内码
	ShowStatus       int      `form:"showStatus"` // 显示状态(1:显示 2:隐藏)
	Status           int      `form:"status"`     // 状态(1:启用 2:禁用)
}

// MenuQueryOptions 查询可选参数项
type MenuQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// MenuQueryResult 查询结果
type MenuQueryResult struct {
	Data       Menus
	PageResult *PaginationResult
}

// Menus 菜单列表
type Menus []*Menu

func (a Menus) Len() int {
	return len(a)
}

func (a Menus) Less(i, j int) bool {
	return a[i].Sequence < a[j].Sequence
}

func (a Menus) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ToMap 转换为键值映射
func (a Menus) ToMap() map[string]*Menu {
	m := make(map[string]*Menu)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

// SplitParentIDs 拆分父级路径的唯一标识列表
func (a Menus) SplitParentIDs() []string {
	idList := make([]string, 0, len(a))
	mIDList := make(map[string]struct{})

	for _, item := range a {
		if _, ok := mIDList[item.ID]; ok || item.ParentPath == "" {
			continue
		}

		for _, pp := range strings.Split(item.ParentPath, "/") {
			if _, ok := mIDList[pp]; ok {
				continue
			}
			idList = append(idList, pp)
			mIDList[pp] = struct{}{}
		}
	}

	return idList
}

// ToTree 转换为菜单树
func (a Menus) ToTree() MenuTrees {
	list := make(MenuTrees, len(a))
	for i, item := range a {
		list[i] = &MenuTree{
			ID:         item.ID,
			Name:       item.Name,
			Icon:       item.Icon,
			RouteName:  item.RouteName,
			RoutePath:  item.RoutePath,
			Component:  item.Component,
			ParentID:   item.ParentID,
			ParentPath: item.ParentPath,
			Sequence:   item.Sequence,
			ShowStatus: item.ShowStatus,
			Status:     item.Status,
			Actions:    item.Actions,
			Resources:  item.Resources,
		}
	}
	return list.ToTree()
}

// FillMenuAction 填充菜单动作列表
func (a Menus) FillMenuAction(mActions map[string]MenuActions) Menus {
	for _, item := range a {
		if v, ok := mActions[item.ID]; ok {
			item.Actions = v
		}
	}
	return a
}

// ----------------------------------------MenuTree--------------------------------------

// MenuTree 菜单树
type MenuTree struct {
	ID         string        `yaml:"-" json:"id"`                                  // 唯一标识
	Name       string        `yaml:"name" json:"name"`                             // 菜单名称
	Icon       string        `yaml:"icon" json:"icon"`                             // 菜单图标
	RouteName  string        `yaml:"routeName,omitempty" json:"routeName"`         // 路由名称
	RoutePath  string        `yaml:"routePath,omitempty" json:"routePath"`         // 路由地址
	Component  string        `yaml:"component,omitempty" json:"component"`         // 组件路径
	ParentID   string        `yaml:"-" json:"parentId"`                            // 父级ID
	ParentPath string        `yaml:"-" json:"parentPath"`                          // 父级路径
	Sequence   int           `yaml:"sequence" json:"sequence"`                     // 排序值
	ShowStatus int           `yaml:"-" json:"showStatus"`                          // 显示状态(1:显示 2:隐藏)
	Status     int           `yaml:"-" json:"status"`                              // 状态(1:启用 2:禁用)
	Actions    MenuActions   `yaml:"actions,omitempty" json:"actions"`             // 动作列表
	Resources  MenuResources `yaml:"resources,omitempty" json:"resources"`         // 资源列表
	Children   *MenuTrees    `yaml:"children,omitempty" json:"children,omitempty"` // 子级树
}

// MenuTrees 菜单树列表
type MenuTrees []*MenuTree

// ToTree 转换为树形结构
func (a MenuTrees) ToTree() MenuTrees {
	mi := make(map[string]*MenuTree)
	for _, item := range a {
		mi[item.ID] = item
	}

	var list MenuTrees
	for _, item := range a {
		if item.ParentID == "" {
			list = append(list, item)
			continue
		}
		if pitem, ok := mi[item.ParentID]; ok {
			if pitem.Children == nil {
				children := MenuTrees{item}
				pitem.Children = &children
				continue
			}
			*pitem.Children = append(*pitem.Children, item)
		}
	}
	return list
}

// ----------------------------------------MenuAction--------------------------------------

// MenuAction 菜单动作对象
type MenuAction struct {
	ID        string              `yaml:"-" json:"id"`                          // 唯一标识
	MenuID    string              `yaml:"-" json:"menuId"`                      // 菜单ID
	Code      string              `yaml:"code" binding:"required" json:"code"`  // 动作编号
	Name      string              `yaml:"name" binding:"required" json:"name"`  // 动作名称
	Resources MenuActionResources `yaml:"resources,omitempty" json:"resources"` // 资源列表
}

// MenuActionQueryParam 查询条件
type MenuActionQueryParam struct {
	PaginationParam
	MenuID string   // 菜单ID
	IDs    []string // 唯一标识列表
}

// MenuActionQueryOptions 查询可选参数项
type MenuActionQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// MenuActionQueryResult 查询结果
type MenuActionQueryResult struct {
	Data       MenuActions
	PageResult *PaginationResult
}

// MenuActions 菜单动作管理列表
type MenuActions []*MenuAction

// ToMap 转换为map
func (a MenuActions) ToMap() map[string]*MenuAction {
	m := make(map[string]*MenuAction)
	for _, item := range a {
		m[item.Code] = item
	}
	return m
}

// FillResources 填充资源数据
func (a MenuActions) FillResources(mResources map[string]MenuActionResources) {
	for i, item := range a {
		a[i].Resources = mResources[item.ID]
	}
}

// ToMenuIDMap 转换为菜单ID映射
func (a MenuActions) ToMenuIDMap() map[string]MenuActions {
	m := make(map[string]MenuActions)
	for _, item := range a {
		m[item.MenuID] = append(m[item.MenuID], item)
	}
	return m
}

// ----------------------------------------MenuActionResource--------------------------------------

// MenuActionResource 菜单动作关联资源对象
type MenuActionResource struct {
	ID          string `yaml:"-" json:"id"`                            // 唯一标识
	ActionID    string `yaml:"-" json:"actionId"`                      // 菜单动作ID
	ResourceID  string `yaml:"-" binding:"required" json:"resourceId"` // 资源ID
	Method      string `yaml:"method" binding:"required" json:"-"`     // 资源请求方式(支持正则)
	Path        string `yaml:"path" binding:"required" json:"-"`       // 资源请求路径（支持/:id匹配）
	Group       string `yaml:"group" binding:"required" json:"-"`      // 资源组
	Description string `yaml:"description" json:"-"`                   // 资源描述
}

// MenuActionResourceQueryParam 查询条件
type MenuActionResourceQueryParam struct {
	PaginationParam
	MenuID  string   // 菜单ID
	MenuIDs []string // 菜单ID列表
}

// MenuActionResourceQueryOptions 查询可选参数项
type MenuActionResourceQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// MenuActionResourceQueryResult 查询结果
type MenuActionResourceQueryResult struct {
	Data       MenuActionResources
	PageResult *PaginationResult
}

// MenuActionResources 菜单动作关联资源管理列表
type MenuActionResources []*MenuActionResource

// ToMap 转换为map
func (a MenuActionResources) ToMap() map[string]*MenuActionResource {
	m := make(map[string]*MenuActionResource)
	for _, item := range a {
		m[item.ResourceID] = item
	}
	return m
}

// ToActionIDMap 转换为动作ID映射
func (a MenuActionResources) ToActionIDMap() map[string]MenuActionResources {
	m := make(map[string]MenuActionResources)
	for _, item := range a {
		m[item.ActionID] = append(m[item.ActionID], item)
	}
	return m
}

// ----------------------------------------MenuResource--------------------------------------

// MenuResource 菜单资源对象
type MenuResource struct {
	ID          string `yaml:"-" json:"id"`                            // 唯一标识
	MenuID      string `yaml:"-" json:"menuId"`                        // 菜单ID
	ResourceID  string `yaml:"-" json:"resourceId" binding:"required"` // 资源ID
	Group       string `yaml:"group" binding:"required" json:"-"`      // 资源组
	Method      string `yaml:"method" binding:"required" json:"-"`     // 资源请求方式(支持正则)
	Path        string `yaml:"path" binding:"required" json:"-"`       // 资源请求路径（支持/:id匹配）
	Description string `yaml:"description" json:"-"`                   // 资源描述
}

func (a *MenuResource) String() string {
	return json.MarshalToString(a)
}

// MenuResourceQueryParam 查询条件
type MenuResourceQueryParam struct {
	PaginationParam
	MenuID  string   // 菜单ID
	MenuIDs []string // 菜单ID列表
}

// MenuResourceQueryOptions 查询可选参数项
type MenuResourceQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// MenuResourceGetOptions Get查询可选参数项
type MenuResourceGetOptions struct {
}

// MenuResourceQueryResult 查询结果
type MenuResourceQueryResult struct {
	Data       MenuResources
	PageResult *PaginationResult
}

// MenuResources 菜单资源列表
type MenuResources []*MenuResource

// ToMap 转换为map
func (a MenuResources) ToMap() map[string]*MenuResource {
	m := make(map[string]*MenuResource)
	for _, item := range a {
		m[item.ResourceID] = item
	}
	return m
}

// ToMenuIDMap 转换为菜单ID映射
func (a MenuResources) ToMenuIDMap() map[string]MenuResources {
	m := make(map[string]MenuResources)
	for _, item := range a {
		m[item.MenuID] = append(m[item.MenuID], item)
	}
	return m
}

// ToMenuIDs 转换为菜单ID列表
func (a MenuResources) ToMenuIDs() []string {
	var idList []string
	m := make(map[string]struct{})

	for _, item := range a {
		if _, ok := m[item.MenuID]; ok {
			continue
		}
		idList = append(idList, item.MenuID)
		m[item.MenuID] = struct{}{}
	}

	return idList
}
