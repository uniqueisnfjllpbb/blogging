package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uniqueisnfjllpbb/blogging/model"
)

func Signup(c *gin.Context) {
	var signupform model.SignupInfo

	fmt.Println(signupform)

}
