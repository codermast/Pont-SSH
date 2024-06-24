package database

import (
	"PontSsh/backend/entity"
	"github.com/google/uuid"
	"log"
)

// SaveSshConnect 保存连接配置
func SaveSshConnect(config entity.SSHConfig) error {

	// 拼接 SQL
	sql := `INSERT INTO sshs (id,server,port,username,password,name,edit) VALUES (?,?,?,?,?,?,?)`

	err := Execute(sql,
		uuid.New().String(),
		config.Server,
		config.Port,
		config.Username,
		config.Password,
		config.Name,
		1)
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
		var editInt int
		err = rows.Scan(&server.Id, &server.Server, &server.Port, &server.Username, &server.Password, &server.Name, &editInt)
		if err != nil {
			log.Printf("Failed to fetch ssh list")
			return nil, err
		}

		if editInt == 1 {
			server.Edit = true
		} else {
			server.Edit = false
		}

		serverList = append(serverList, server)
	}

	return serverList, err
}

// UpdateConnection 更新连接
func UpdateConnection(sshConfig entity.SSHConfig) error {
	sql := "UPDATE sshs SET server = ? , port = ?,username = ?,password = ?, name = ?, edit = ? WHERE id = ?"
	var editInt int

	if sshConfig.Edit == true {
		editInt = 1
	} else {
		editInt = 0
	}
	err := Execute(sql, sshConfig.Server, sshConfig.Port, sshConfig.Username, sshConfig.Password, sshConfig.Name, editInt, sshConfig.Id)
	return err
}

// SearchConnection 搜索连接
func SearchConnection(keyword string) ([]entity.SSHConfig, error) {
	sql := "SELECT * FROM sshs WHERE name LIKE '%" + keyword +
		"%' OR server LIKE '%" + keyword +
		"%' OR username LIKE '%" + keyword + "%'"

	rows, err := Query(sql)
	if err != nil {
		log.Printf("Failed to Query ssh list")
		return nil, err
	}
	var serverList []entity.SSHConfig
	for rows.Next() {
		var server entity.SSHConfig
		var editInt int

		err = rows.Scan(&server.Id, &server.Server, &server.Port, &server.Username, &server.Password, &server.Name, &editInt)

		if err != nil {
			log.Printf("Failed to fetch ssh list")
			return nil, err
		}

		if editInt == 1 {
			server.Edit = true
		} else {
			server.Edit = false
		}
		serverList = append(serverList, server)
	}
	return serverList, err
}
