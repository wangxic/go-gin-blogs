package controllers

import (
	"../../gin/models"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
}

func (_ *Login) Index(ctx *gin.Context) {

	title := "heihei"

	if ctx.Request.Method == "POST" {

		password := ctx.PostForm("password")
		account := ctx.PostForm("account")
		// 	if len(password) == 0 || len(account) == 0 {

		// 	}
		// 	//验证登录
		// 	dd := memberModel.First(account)
		// 	//
		// 	// 数据
		memberModel := new(models.Member)
		Member := memberModel.First(account)
		//查询不到结果
		if Member.Mid == 0 {
			ctx.JSON(200, gin.H{"result": 0, "msg": "没有找到这个用户"})
			return
		}

		s := password + Member.Salt
		//产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
		h := sha1.New() // md5加密类似md5.New()
		//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
		h.Write([]byte(s))
		//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来对现有的字符切片追加额外的字节切片：一般不需要要。
		bs := h.Sum(nil)
		//SHA1 值经常以 16 进制输出，使用%x 来将散列结果格式化为 16 进制字符串。
		encodedStr := hex.EncodeToString(bs)

		//如果需要对另一个字符串加密，要么重新生成一个新的散列，要么一定要调用h.Reset()方法，不然生成的加密字符串会是拼接第一个字符串之后进行加密
		h.Reset() //重要！！！

		if encodedStr != Member.Password {
			ctx.JSON(200, gin.H{"result": 0, "msg": "密码错误"})
			return
		}
		//设置session
		session := sessions.Default(ctx)
		option := sessions.Options{MaxAge: 30 * 600}
		session.Options(option)
		session.Set("sessionid", Member.Mid)
		session.Save()
		// 登录成功
		ctx.JSON(200, gin.H{"result": 1, "msg": "登录成功", "data": Member})
		return
	}

	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "login.html", gin.H{"title": title, "data": "我来了"})
	}

}
