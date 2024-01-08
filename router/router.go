package router

import (
	"github.com/gin-gonic/gin"
	"github.com/uniqueisnfjllpbb/blogging/controller"
	"github.com/uniqueisnfjllpbb/blogging/handler"
)

func Routes(rtr gin.IRouter) {

	gin.New()

	rtr.POST("/register", controller.Register)

	//rtr.POST("/signup", SignUp)
	//rtr.POST("login", SignUp)
	//rtr.POST("/logout", LogOut)

	//後々メソッドは消したい（動作確認後）

	//編集
	apis := rtr.Group("/api")
	apis.POST("/createapost", handler.AddAPost)
	apis.GET("/showallposts", handler.ShowAllPosts)
	apis.DELETE("/deleteblog/:id", handler.DeleteAPost)
	apis.GET("/showapost/:id", handler.ShowAPost)
	apis.PUT("/rewrite/:id", handler.RewriteAPost)

	//クライアント側
	blog := rtr.Group("/blog")
	blog.GET("showallposts", handler.ShowAllPosts)
	blog.GET("addapost/:id", handler.AddAPost)

}
