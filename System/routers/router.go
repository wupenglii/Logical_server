package routers

import (
	"Logical_server/System/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//首页&分类&详情

	beego.Router("/", &controllers.HomeController{}, "get:Index2")
	beego.Router("/2", &controllers.HomeController{}, "get:Index2")

	//首页&分类&详情
	// beego.Router("/", &controllers.HomeController{}, "get:Index")               //首页
	// beego.Router("/explore", &controllers.ExploreController{}, "get:Index")     //分类
	// beego.Router("/books/:key", &controllers.DocumentController{}, "get:Index") //详情

	// //货物编辑
	// beego.Router("/api/:key/content/?:id", &controllers.DocumentController{}, "*:Content") //保存文档内容
	// beego.Router("/api/:key/edit/?:id", &controllers.DocumentController{}, "*:Edit")       //文档编辑
	// beego.Router("/api/upload", &controllers.DocumentController{}, "post:Upload")          //上传图片
	// beego.Router("/read/:key/create", &controllers.DocumentController{}, "post:Create")
	// beego.Router("/read/:key/delete", &controllers.DocumentController{}, "post:Delete")

	// //搜索
	// beego.Router("/search", &controllers.SearchController{}, "get:Search")
	// beego.Router("/search/result", &controllers.SearchController{}, "get:Result")

	//login
	beego.Router("/login", &controllers.AccountController{}, "*:Login")
	beego.Router("/regist", &controllers.AccountController{}, "*:Regist")
	beego.Router("/logout", &controllers.AccountController{}, "*:Logout")
	beego.Router("/doregist", &controllers.AccountController{}, "post:DoRegist")

	// //仓库货物管理
	// beego.Router("/book", &controllers.BookController{}, "*:Index")                         //我的货物
	// beego.Router("/book/create", &controllers.BookController{}, "post:Create")              //创建货物
	// beego.Router("/book/:key/setting", &controllers.BookController{}, "*:Setting")          //货物设置
	// beego.Router("/book/setting/upload", &controllers.BookController{}, "post:UploadCover") //货物封面
	// beego.Router("/book/setting/save", &controllers.BookController{}, "post:SaveBook")      //保存
	// beego.Router("/book/:key/release", &controllers.BookController{}, "post:Release")       //发布货物

	// //个人中心
	// beego.Router("/user/:username", &controllers.UserController{}, "get:Index")                 //分享
	// beego.Router("/user/:username/collection", &controllers.UserController{}, "get:Collection") //收藏
	// beego.Router("/user/:username/follow", &controllers.UserController{}, "get:Follow")         //关注
	// beego.Router("/user/:username/fans", &controllers.UserController{}, "get:Fans")             //粉丝
	// beego.Router("/follow/:uid", &controllers.BaseController{}, "get:SetFollow")                //关注或取消关注
	// beego.Router("/book/score/:id", &controllers.BookController{}, "*:Score")                   //分享
	// beego.Router("/book/comment/:id", &controllers.BookController{}, "post:Comment")            //评论

	// //个人设置模块
	// beego.Router("/setting", &controllers.SettingController{}, "*:Index")
	// beego.Router("/setting/upload", &controllers.SettingController{}, "*:Upload")

	// //管理后台
	// beego.Router("/manager/category", &controllers.ManagerController{}, "post,get:Category")
	// beego.Router("/manager/update-cate", &controllers.ManagerController{}, "get:UpdateCate")
	// beego.Router("/manager/del-cate", &controllers.ManagerController{}, "get:DelCate")
	// beego.Router("/manager/icon-cate", &controllers.ManagerController{}, "post:UpdateCateIcon")

}
