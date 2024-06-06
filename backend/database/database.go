package database

import "database/sql"

// 数据库连接
var db *sql.DB

// InitDatabase 初始化数据库连接
func InitDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "./pontssh.db")
	return err
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
func Query(query string) (*sql.Rows, error) {
	rows, err := db.Query(query)

	return rows, err
}

// Execute 执行其他语句
func Execute(query string) error {
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}