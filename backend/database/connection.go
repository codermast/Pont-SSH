package database

import (
	"PontSsh/backend/entity"
	"github.com/google/uuid"
	"log"
)

// SaveSshConnect 保存连接配置
func SaveSshConnect(config entity.SSHConfig) error {

	// 拼接 SQL
	sql := `INSERT INTO sshs (id,host,port,username,password,comment) VALUES (?,?,?,?,?,?)`

	err := Execute(sql,
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

// GetServerList 查询服务器列表
func GetServerList() ([]entity.SSHConfig, error) {
	// 1. 拼接查询 SQL
	sql := "SELECT * FROM sshs"

	// 2. 执行 SQL
	rows, err := Query(sql)
	if err != nil {
		log.Printf("Failed to Query ssh list")
		return nil, err
	}

	var serverList []entity.SSHConfig

	for rows.Next() {
		var server entity.SSHConfig
		var id string
		err = rows.Scan(&id, &server.Server, &server.Port, &server.Username, &server.Password, &server.Name)
		if err != nil {
			log.Printf("Failed to fetch ssh list")
			return nil, err
		}
		log.Printf("%d", id)

		server.Edit = false

		serverList = append(serverList, server)
	}

	return serverList, err
}
