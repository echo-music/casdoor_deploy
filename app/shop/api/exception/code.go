package exception

const (
	CodeOK               = 1000
	CodeInvalidParameter = 1400
	CodeUnauthorized     = 1401
	CodeFailure          = 1402
	CodeForbidden        = 1403
	CodeServerError      = 1500
	CodeServerTimeout    = 1408
)

func CodeText(code int) string {
	switch code {
	case CodeOK:
		return "OK"
	case CodeInvalidParameter:
		return "无效的参数"
	case CodeFailure:
		return "操作失败"
	case CodeServerError:
		return "系统内部错误"
	case CodeServerTimeout:
		return "服务超时"
	case CodeUnauthorized:
		return "用户未授权"
	default:
		return "未知的错误"
	}
}
