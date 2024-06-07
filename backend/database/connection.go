package database

import (
	"PontSsh/backend/entity"
	"github.com/google/uuid"
)

// SaveSshConnect 保存连接配置
func SaveSshConnect(config entity.SSHConfig) error {
	err := InitDatabase()
	// 数据库初始化异常
	if err != nil {
		return err
	}

	// 拼接 SQL
	sql := `INSERT INTO sshs (id,server,port,username,password,keypath,dir) VALUES (?,?,?,?,?,?,?)`

	err = Execute(sql, uuid.New().String(), config.Server, config.Port, config.Username, config.Password, config.KeyPath, "/")

	return err
}
