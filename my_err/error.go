package my_err

type MyErr int

const (
	ErrInputForm MyErr = iota + 40300001
	ErrUserNotFound
	ErrUserAlreadyExists
	ErrInvalidCredentials
)

const (
	ErrDataBaseFail MyErr = 50000001 + iota
	ErrServer
)

func (e MyErr) Code() int {
	return int(e)
}
func (e MyErr) Error() string {
	switch e {
	case ErrInputForm:
		return "输入格式错误"
	case ErrUserNotFound:
		return "用户不存在"
	case ErrUserAlreadyExists:
		return "用户已存在"
	case ErrInvalidCredentials:
		return "用户名或密码错误"
	case ErrDataBaseFail:
		return "数据库错误"
	case ErrServer:
		return "服务器错误"
	default:
		return "未知错误"
	}
}
