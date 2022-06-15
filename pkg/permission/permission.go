package permission

type Permission interface {
	// Policy有关
	// CheckPermission 检测权限
	CheckPermission(roleId, resource, action string) (bool, error)
	// RemovePermission 删除权限
	RemovePermission(presetId string) error
	// AddPermissions 添加权限
	AddPermissions(presetId, effect string, actions, resources []string) error
	// UpdatePermission 更新权限
	UpdatePermission(presetId, effect string, actions, resources []string) error

	// Group有关
	// AddPermissionGroup 添加组关系
	AddPermissionGroup(roleId, presetId string) error
	// DeletePermissionGroup 删除权限组
	DeletePermissionGroup(roleId, presetId string) error
	// DeleteAppPermissionGroup 删除特定一组权限组
	DeleteAppPermissionGroup(roleId string) error
}