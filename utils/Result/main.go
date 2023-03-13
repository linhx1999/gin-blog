package Result

import "net/http"

var errCodeMsg = map[string]string{
	"00000": "一切ok", // 正确执行后的返回

	"A0001": "用户端错误",  // 一级宏观错误码
	"A0100": "用户注册错误", // 二级宏观错误码
	"A0200": "用户登录异常",
	"A0300": "访问权限异常",

	"A0400": "用户请求参数错误",
	"A0402": "无效的用户输入",
	"A0420": "请求参数值超出允许的范围",

	"A0500": "用户请求服务异常",

	"A0600": "用户资源异常 ",
	"A0605": "用户配额已用光 ", // 蚂蚁森林浇水数或每天抽奖数

	"A0700":  "用户上传文件异常",
	"A0800":  "用户当前版本异常 ",
	"A0900 ": "用户隐私未授权",

	"B0001":  "系统执行出错",
	"B0100 ": "系统执行超时 ",
	"B0200":  "系统容灾功能被触发",
	"B0300":  "系统资源异常 ",

	"C0001": "调用第三方服务出错",
	"C0100": "中间件服务出错 ",
	"C0110": "RPC 服务出错  ",
	"C0200": "第三方系统执行超时 ",
	"C0400": "第三方容灾系统被触发 ",
	"C0500": "通知服务出错 ",
}

type result struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrData struct {
	Code    string `json:"code"`
	ErrMsg  string `json:"errMsg"`
	UserMsg string `json:"userMsg"`
}

func NewSuccess(msg string, data []interface{}) *result {
	return &result{
		Status:  http.StatusOK,
		Message: msg,
		Data:    data,
	}
}

func NewFail(status int, msg string, errorCode string, userMsg string) *result {
	return &result{
		Status:  status,
		Message: msg,
		Data: []any{ErrData{
			errorCode,
			errCodeMsg[errorCode],
			userMsg,
		}},
	}
}
