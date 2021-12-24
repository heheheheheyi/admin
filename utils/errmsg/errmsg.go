package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//数据库
	ErrorSql = 700

	ErrorUpload = 800

	ErrorTokenExist     = 901
	ErrorTokenWrong     = 903
	ErrorTokenRuntime   = 904
	ErrorTokenTypeWrong = 905

	ErrorUsernameUsed  = 1001
	ErrorUserNotExist  = 1002
	ErrorPasswordWrong = 1003
	ErrorUserNoRole    = 1004
	ErrorUserHasBollow = 1005

	ErrorCatenameUsed = 2001
	ErrorCateHasRes   = 2002

	ErrorResnameUsed  = 3001
	ErrorResHasBorrow = 3002
)

var codemsg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	ErrorSql: "读取数据失败",

	ErrorUpload: "上传失败",

	ErrorTokenExist:     "Token不存在",
	ErrorTokenWrong:     "Token错误",
	ErrorTokenRuntime:   "Token过期",
	ErrorTokenTypeWrong: "Token类型错误",

	ErrorUsernameUsed:  "用户名已存在",
	ErrorUserNotExist:  "用户不存在",
	ErrorPasswordWrong: "密码错误",
	ErrorUserNoRole:    "用户无权限",
	ErrorUserHasBollow: "用户存在未归还资源",

	ErrorCatenameUsed: "分类名已存在",
	ErrorCateHasRes:   "分类下存在资源",

	ErrorResnameUsed:  "资源名已存在",
	ErrorResHasBorrow: "资源已被借走",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
