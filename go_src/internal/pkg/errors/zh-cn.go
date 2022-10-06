package errors

var zhCNText = map[int]string{
	SUCCESS:            "OK",
	FAILURE:            "FAIL",
	NotFound:           "资源不存在",
	ServerError:        "服务器内部错误",
	TooManyRequests:    "请求过多",
	InvalidParameter:   "参数错误",
	UserDoesNotExist:   "用户不存在",
	AuthorizationError: "权限错误",
	RBACError:          "暂无访问权限",
	UserISExist:        "用户已经存在",
	BLOGIDNOTFOUND:     "该博客不存在",
	NOTOKEN:            "没有token值，你没有权限访问此路径",
	TOKENERROR:         "错误的token值",
}
