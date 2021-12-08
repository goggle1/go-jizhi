package common

const (
	ResultSuccess          = 0
	ResultErrorVerifi      = -iota //验证错误
	ResultErrorParameter           //参数错误
	ResultErrorPermissions         //权限错误
	ResultErrorSystem              //系统错误
	ResultErrorAction              //操作错误
	ResultErrorSession
)

const (
	SecretKey    = "3fcb8e8216fab5%^(3074926"
	SecretVector = "34814563"
)
