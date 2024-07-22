package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-todo-list/common"
)

func Recovery() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if err, ok := err.(error); ok {
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrInvalidRequest(err))
				}
				panic(err)
			}
		}()
		ctx.Next()
	}
}
