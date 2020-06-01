package sysinit

import (
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

func sysinit() {
	upload := filepath.Join("./", "uploads")
	beego.BConfig.WebConfig.StaticDir["/uploads"] = upload

	//注册前端使用函数
	registerFunctions()
}

func registerFunctions() {
	beego.AddFuncMap("cdnjs", func(p string) string {
		cdn := beego.AppConfig.DefaultString("cdnjs", "")
		if strings.HasPrefix(p, "/") && strings.HasSuffix(cdn, "/") {
			return cdn + string(p[1:])
		}
		return ""
	})

}
