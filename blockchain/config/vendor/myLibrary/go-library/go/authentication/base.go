package authentication

// 数据权限
type AuthTypeData uint64

// 操作权限
type AuthTypeOperation uint64

type AuthValue uint64

type AuthType int8

// 操作权限定义:
const (
)

// 模块定义
const (


)

// 操作&数据权限值定义
// 对于重复的操作权限不使用公共字段,而是单独定义,若使用公共字段会导致无法区分数据权限和操作权限
const (
)

// 产品常量信息定义
const (
)
