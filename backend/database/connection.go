package database

import (
	"PontSsh/backend/entity"
	"github.com/google/uuid"
	"log"
)

// SaveSshConnect 保存连接配置
func SaveSshConnect(config entity.SSHConfig) error {
	err := InitDatabase()
	// 数据库初始化异常
	if err != nil {
		log.Println("Failed to connect to database")
		return err
	}

	// 拼接 SQL
	sql := `INSERT INTO sshs (id,host,port,username,password,comment) VALUES (?,?,?,?,?,?)`

	err = Execute(sql,
		uuid.New().String(),
		config.Server,
		config.Port,
		config.Username,
		config.Password,
		"123")
	if err != nil {
		log.Println("Failed to save ssh connection")
		return err
	}
	return nil
}
