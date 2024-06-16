package database

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

// 数据库连接
var db *sql.DB

// InitDatabase 初始化数据库连接
func InitDatabase() error {
	var err error
	db, err = sql.Open("sqlite", "./pontssh.db")

	if err != nil {
		return err
	}

	// Ping 数据库来验证连接
	if err = db.Ping(); err != nil {
		return err
	}

	sql := "SELECT * FROM sshs"

	rows, err := db.Query(sql)

	if err != nil {
		return err
	}

	i := 1
	// 遍历结果集
	for rows.Next() {
		var id, host, username, password, comment string
		var port int
		if err := rows.Scan(&id, &host, &port, &username, &password, &comment); err != nil {
			log.Printf("Scan error: %v", err)
			return err
		}
		log.Printf("%d - Row: id=%s, host=%s, port=%d, username=%s, password=%s, comment=%s", i, id, host, port, username, password, comment)
		i++
	}

	// 检查遍历过程中是否有错误
	if err := rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return err
	}

	defer rows.Close()

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
func Query(query string) (*sql.Rows, error) {
	rows, err := db.Query(query)

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
