package database

import (
	"PontSsh/backend/constant"
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

// 数据库连接
var db *sql.DB

// InitDatabase 初始化数据库连接
func InitDatabase() error {
	// 连接数据库
	dbOpen, err := sql.Open("sqlite", constant.DatabaseFilePath)

	if err != nil {
		return err
	}

	db = dbOpen

	// Ping 数据库来验证连接
	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

// CloseDatabase 关闭数据库连接
func CloseDatabase() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

// Query 执行查询语句
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	return rows, err
}

// Execute 执行其他语句
func Execute(sql string, args ...interface{}) error {
	_, err := db.Exec(sql, args...)

	if err != nil {
		log.Println("sql 执行异常：", err)
		return err
	}
	return nil
}
