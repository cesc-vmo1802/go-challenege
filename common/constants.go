package common

import "time"

const (
	RequesterKey = "requester"

	DefaultDatabase = "onqlavelabs"

	UserRoleCacheKey       = "user:%v:roles"
	UserPermissionCacheKey = "user:%v:permissions"

	DefaultCacheExpiration = time.Hour * 24
)
