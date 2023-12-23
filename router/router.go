package router

import (
	"github.com/gin-gonic/gin"
	"github.com/uniqueisnfjllpbb/blogging/handler"
)

func Routes(r gin.IRouter) {

	//rtr := gin.New()

	//rtr.POST("/signup", SignUp)
	//rtr.POST("login", SignUp)
	//rtr.POST("/logout", LogOut)

	routeblog := r.Group("/blog")

	routeblog.GET("/showallposts", handler.ShowAllPosts)

}
