package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数 （参数校验）
func SignUpHandler(c *gin.Context) {
	//获得参数和校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("signup failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//业务响应
	logic.SignUp(p)
}
