package common

import "github.com/astaxie/beego"

//session
const SessionName = "__system_session__"

//正则表达式
const RegexpEmail = `^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`

//默认PageSize
const PageSize = 20
const RollPage = 4

const WorkingDirectory = "./"

//用户权限
const (
	//超级管理员
	MemberSuperRole = 0
	//普通管理员
	MemberAdminRole = 1
	//普通用户
	MemberGeneralRole = 2
)

func Role(role int) string {
	if role == MemberSuperRole {
		return "超级管理员"
	} else if role == MemberAdminRole {
		return "管理员"
	} else if role == MemberGeneralRole {
		return "普通用户"
	} else {
		return ""
	}
}

//app_key
func AppKey() string {
	return beego.AppConfig.DefaultString("app_key", "godoc")
}

//默认头像
func DefaultAvatar() string {
	return beego.AppConfig.DefaultString("avatar", "/static/images/headimgurl.jpg")
}
