package err

type code uint32

const (
	// 创建资源失败
	CodeCreateFialed code = 1001001
	// 记录不存在
	CodeItemNotFound code = 1001002
	// 更新记录失败
	CodeUpdateFialed code = 1001003
	// 删除记录失败
	CodeDeleteFialed code = 1001004

	// 记录已过期
	CodeItemExpired code = 2001001

	// 资源繁忙
	CodeIsBusy code = 3001001
)

func (e *errs) Code(code uint32) Err {
	e.code = code
	return e
}
