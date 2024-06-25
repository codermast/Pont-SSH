package service

import (
	"PontSsh/assets"
	"PontSsh/backend/constant"
	"PontSsh/backend/database"
	"context"
	"database/sql"
	_ "embed"
	"log"
	"os"
	"path/filepath"
)

type Project struct {
	ctx context.Context
}

func NewProject() *Project {
	return &Project{}
}

// Startup 启动方法
func (p *Project) Startup(ctx context.Context) {
	// 加载数据库信息
	loadDatabase()
}

// Shutdown 销毁方法
func (p *Project) Shutdown(ctx context.Context) {
	err := database.CloseDatabase()

	if err != nil {
		log.Println("关闭数据库连接失败！")
	}
}

// loadDatabase 加载数据库信息
func loadDatabase() {
	filePath := constant.DatabaseFilePath

	log.Println("文件路径：" + filePath)

	// 0. 判断是否需要初始化
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		log.Println("数据库文件不存在，则需要进行初始化")

		// 1. 创建数据库文件

		err := os.MkdirAll(filepath.Dir(filePath), 0700)

		if err != nil {
			log.Printf("创建目录异常：%v", err)
		} else {
			log.Println("创建目录成功：" + filepath.Dir(filePath))
		}

		_, err = os.Create(filePath)

		if err != nil {
			log.Printf("创建文件异常：%v", err)
		} else {
			log.Println("创建文件成功")
		}

		// 2. 连接数据库
		// 连接数据库
		dbOpen, err := sql.Open("sqlite", constant.DatabaseFilePath)

		if err != nil {
			log.Printf("数据库连接异常：%v", err)
		} else {
			log.Println("数据库连接成功")
		}

		// 3. 执行数据库初始化
		_, err = dbOpen.Exec(assets.InitSql)

		if err != nil {
			log.Printf("数据库初始化失败：%v", err)
		} else {
			log.Println("数据库初始化完毕！")
		}
	} else {
		log.Println("数据库已经存在，则不必初始化！")
	}
}
