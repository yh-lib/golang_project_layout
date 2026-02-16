// 项目的总入口
package main

import (
	_ "golang_project_layout/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 开始加载程序配置
	// 2. 配置 gin
	r := gin.Default()
	r.Run()
}
