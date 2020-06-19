package controllers

import (
	// "gin-learning/models"

	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "strconv"
)

// 什么一个空的结构体
type Ui struct {
}

func (_ *Ui) Index(ctx *gin.Context) {
	// 声明一个map数组
	title := make(map[string]string)

	// articleModel := new(models.Articles)
	// list := articleModel.List()
	// 类型转换 string转float
	// 在转换时浮点型的时候需要特别注意精度问题
	// 在float32 下值等于 10.010000228881836
	// 在float64 下等于10.01
	// int_nut := "10.01"
	// folatd, _ := strconv.ParseFloat(int_nut, 64)
	// fmt.Println(folatd)

	// 转字符串
	// 格式fmt是'b'（-ddddp±ddd，二进制指数），
	// 'e'（-d.dddde±dd，十进制指数），
	// 'E'（-d.ddddE±dd，十进制指数）之一），
	// 'f'（-ddd.dddd，无指数），
	// 'g'（大指数为'e'，否则为'f'），
	// 'G'（大指数为'E'，否则为'f'），
	// ' x'（-0xd.ddddp±ddd，十六进制分数和二进制指数）或'
	// X'（-0Xd.ddddP±ddd，十六进制分数和二进制指数）
	// FormatFloat(float float64,'f','保留的精度'，float64);
	// 特殊精度-1使用必要的最小位数
	// 64 位宽
	// str := strconv.FormatFloat(folatd, 'f', 2, 64)
	// fmt.Println(str)

	// // 字符转int 10进
	// // 64位宽
	// str_int := "10"
	// int_1, _ := strconv.ParseInt(str_int, 10, 64)
	// fmt.Println(int_1)
	// // int转字符串
	// // FormatInt(int64 类型的int，10进制)
	// fmt.Println(strconv.FormatInt(int_1, 10))

	// // map数组赋值
	title["title"] = "测试"

	// for key, v := range title {
	// 	fmt.Printf("map[%v]=%v\n", key, v)
	// }

	// fmt.Println(title)
	// articleModel := new(models.Articles)
	// list := articleModel.List()
	// 模板输出
	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": title, "data": "我来了"})
}

func (_ *Ui) Landing(ctx *gin.Context) {

	// 模板输出
	ctx.HTML(http.StatusOK, "landing.html", gin.H{})
}

func (_ *Ui) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

// func (_ *Articles) Edit(ctx *gin.Context) {
//     id, err := strconv.Atoi(ctx.Param("id"))
//     if err != nil {
//         ctx.Redirect(http.StatusFound, "/articles")
//         return
//     }
//     articleModel := new(models.Articles)
//     article := articleModel.First(id)
//     ctx.HTML(http.StatusOK, "articles/create-edit.html", gin.H{
//         "article": article,
//     })
// }

// func (_ *Articles) Store(ctx *gin.Context) {
//     id, _ := strconv.Atoi(ctx.PostForm("id"))
//     title := ctx.PostForm("title")
//     author := ctx.PostForm("author")
//     content := ctx.PostForm("content")
//     articleModel := new(models.Articles)
//     if id == 0 {
//         articleModel.Insert(title, author, content)
//     } else {
//         articleModel.Edit(id, title, author, content)
//     }

//     ctx.Redirect(http.StatusFound, "/articles")
// }

// func (_ *Articles) Del(ctx *gin.Context) {
//     id, err := strconv.Atoi(ctx.Param("id"))
//     if err != nil {
//         ctx.Redirect(http.StatusFound, "/articles")
//         return
//     }
//     articleModel := new(models.Articles)
//     articleModel.Del(id)
//     ctx.Redirect(http.StatusFound, "/articles")
// }
// ctx *gin.Context  通过路由处传 &代表指针 *代表内存地址
func (_ *Ui) test(ctx *gin.Context) int {
	// 接受post参数
	// title := ctx.PostForm("title")
	return 1
}
