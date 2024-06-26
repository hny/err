package err

// api
const (
	PARAMS_INVALIDED_CODE = 10000 + iota
	PARAMS_REQUIRED_CODE
	EFFECT_ROWS_0_CODE
	REPEATED_REQUEST_CODE
	REPEATED_CHECK_CODE
	USERNAME_OR_PASSWORD_CODE
	INVALIDED_TOKEN_CODE
	ONLY_FOR_ADMIN_CODE
	SYSTEM_BUSY_CODE
	SYSTEM_INNER_CODE
)
const (
	PARAMS_INVALIDED_ZN     = "参数有误，请检查"
	PARAMS_REQUIRED_ZN      = "必填参数不可为空"
	EFFECT_ROWS_0_ZN        = "操作影响数为 0"
	REPEATED_REQUEST_ZN     = "请勿重复请求"
	REPEATED_CHECK_ZN       = "请求去重校验失败"
	USERNAME_OR_PASSWORD_ZN = "用户名或密码错误"
	INVALIDED_TOKEN_ZN      = "Token无效"
	ONLY_FOR_ADMIN_ZN       = "仅支持系统管理员"
	SYSTEM_BUSY_ZN          = "系统繁忙请稍后"
	SYSTEM_INNER_ZN         = "服务异常请稍后"
)

// data
const (
	ITEM_NOT_EXIST_CODE = 10000 + iota
	ITEM_EXPIRED_CODE
	ITEM_ALREADY_EXIST_CODE
	CREATE_FIALED_CODE
	UPDATE_FIALED_CODE
	DELETE_FIALED_CODE
)
const (
	ITEM_NOT_EXIST_ZN = "记录不存在"
	ITEM_EXPIRED_ZN   = "记录已过期"
	ITEM_REPEATED_ZN  = "记录已存在"
	CREATE_FIALED_ZN  = "创建资源失败"
	UPDATE_FIALED_ZN  = "更新记录失败"
	DELETE_FIALED_ZN  = "删除记录失败"
)
