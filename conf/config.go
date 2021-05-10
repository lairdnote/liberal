package conf

import (
	"github.com/joho/godotenv"
	"github.com/lairdnote/liberal/models"
	"os"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()
	// 连接数据库
	models.Database(os.Getenv("POSTGRES_DSN"))
}

