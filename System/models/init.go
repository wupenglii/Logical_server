package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(
		new(Member)
	)
}

func TNMembers() string {
	return "md_members"
}
