package server

import (
	"github.com/gin-gonic/gin"
	"github.com/uniqueisnfjllpbb/blogging/handler"
)

func Routes(r gin.IRouter) {
	routeblog := r.Group("/blog")

	routeblog.GET("/showallposts", handler.ShowAllPosts)

}