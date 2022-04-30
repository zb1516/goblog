package router

import (
	"github.com/julienschmidt/httprouter"
	"goblog/controller"
	"goblog/controller/admin"
	"goblog/middleware"
	"net/http"

	//"goblog/middleware"
	//"net/http"
)

func RegisterRouter() *httprouter.Router{
	router :=httprouter.New()
	router.GET("/",controller.Index)
	router.GET("/tag", controller.TagList)
	router.GET("/category", controller.CategoryList)
	router.GET("/about", controller.About)
	router.GET("/post/:id", controller.Post)
	router.Handler("POST","/love",middleware.ResponseHandler(controller.Love))
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.GET("/admin",admin.Index)
	router.GET("/admin/post-list",admin.PostList)
	router.GET("/admin/post-add",admin.PostAdd)
	router.Handler("POST","/admin/post-save",middleware.ResponseHandler(admin.PostSave))
	router.GET("/admin/post-edit",admin.PostEdit)
	router.Handler("POST","/admin/post-del",middleware.ResponseHandler(admin.PostDel))
	router.Handler("POST","/admin/post-edit-save",middleware.ResponseHandler(admin.PostEditSave))

	router.GET("/admin/tag-list",admin.TagList)
	router.GET("/admin/tag-add",admin.TagAdd)
	router.Handler("POST","/admin/tag-save",middleware.ResponseHandler(admin.TagSave))
	router.GET("/admin/tag-edit",admin.TagEdit)
	router.Handler("POST","/admin/tag-del",middleware.ResponseHandler(admin.TagDel))
	router.Handler("POST","/admin/tag-edit-save",middleware.ResponseHandler(admin.TagEditSave))

	router.GET("/admin/category-list",admin.CategoryList)
	router.GET("/admin/category-add",admin.CategoryAdd)
	router.GET("/admin/category-edit",admin.CategoryEdit)
	router.Handler("POST","/admin/category-del",middleware.ResponseHandler(admin.CategoryDel))
	router.Handler("POST","/admin/category-edit-save",middleware.ResponseHandler(admin.CategoryEditSave))
	router.Handler("POST","/admin/category-save",middleware.ResponseHandler(admin.CategorySave))

	router.GET("/admin/user-list",admin.UserList)
	router.GET("/admin/user-add",admin.UserAdd)
	router.Handler("POST","/admin/user-save",middleware.ResponseHandler(admin.UserSave))
	router.GET("/admin/user-edit",admin.UserEdit)
	router.Handler("POST","/admin/user-del",middleware.ResponseHandler(admin.UserDel))
	router.Handler("POST","/admin/user-edit-save",middleware.ResponseHandler(admin.UserEditSave))

	router.GET("/admin/login",admin.Login)
	router.GET("/admin/logout",admin.Logout)
	router.Handler("POST","/admin/logins",middleware.ResponseHandler(admin.Logins))
	router.GET("/admin/config",admin.Config)
	router.Handler("POST","/admin/config-save",middleware.ResponseHandler(admin.ConfigSave))
	router.Handler("POST","/upload",middleware.ResponseHandler(controller.Upload))	//统一上传接口
	router.GET("/showpic",controller.ShowPicHandle)
	return router
}