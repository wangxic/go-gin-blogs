package middlewares

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func verification(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	fmt.Println(session.Get("boke"))
}
