package conf

// 配置读取
import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var cfg *ini.File

func GetConfig(key string) string {
	cfg, err := ini.Load("../gin/conf/database.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return cfg.Section("develop").Key(key).String()

}
