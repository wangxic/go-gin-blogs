package routers

import (
    "../../gin/controllers"
    // "../../gin/middlewares"
    "fmt"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Register() *gin.Engine {
    r := gin.New()
    r.Use(gin.Recovery())
    //初始化session
    store := cookie.NewStore([]byte("secret"))

    r.Use(sessions.Sessions("mysession", store))
    //加载controllers 包中的Ui文件
    Ui := new(controllers.Ui)
    // 加载中间见

    // 设置默认页
    r.GET("/", Ui.Index)
    // V := r.Group("/"){
    //     V.GET("/", Ui.Index)
    // }
    // 分组路由
    v1 := r.Group("/ui", verification)
    {
        v1.GET("/", Ui.Index)
        v1.GET("/index", Ui.Index)
        v1.GET("/landing", Ui.Landing)
        v1.GET("/login", Ui.Login)
        v1.POST("/login", Ui.Login)
        // v1.GET("/article/create", articles.Create)
        // v1.GET("/article/edit/:id", articles.Edit)
        // v1.GET("/article/del/:id", articles.Del)
        // v1.POST("/article/store", articles.Store)
    }
    Login := new(controllers.Login)

    v2 := r.Group("/login")
    {
        v2.GET("/", Login.Index)
        v2.POST("/", Login.Index)
    }
    return r
}

func verification(ctx *gin.Context) {
    session := sessions.Default(ctx)
    sessionVal := session.Get("sessiondid")

    fmt.Println(ctx.Request.Host)
    if sessionVal == nil {
        ctx.Redirect(http.StatusMovedPermanently, "/login")
    }
}
