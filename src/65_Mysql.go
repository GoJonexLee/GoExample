package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"	// mysql驱动包,并且只需要加载mysql驱动包的init()函数
	"fmt"
	"time"
)

func main() {

	// 数据库的链接字符串格式
	format := "%s:%s@tcp(%s:%d)/%s?charset=utf8"
	dbInfo := fmt.Sprintf(format, "username", "password", "localhost", 8080, "dbName")

	database, _ := sql.Open("mysql", dbInfo)	// 第一个参数是驱动的名称，这里表明是mysql驱动

	database.SetMaxOpenConns(10)	// 设置最大连接数
	database.SetMaxIdleConns(10)	// 设置最大空闲数
	database.SetConnMaxLifetime(10 * time.Second) // 设置链接超时数


	//create、update一般直接使用exec()方法
	result, _ := database.Exec("update table set columnValu = ? where id = ?", "value", "id")

	// 单行查询
	row := database.QueryRow("select id, name from table where id ?", "id")
	var id, name string
	row.Scan(&id, &name)	// 将查询结果赋给id和name，这里一定要取变量的地址


	// RowsAffected():一共更新多少条数据；  LastInsertId():最后一次更新数据的id
	fmt.Println(result.RowsAffected())
	fmt.Println(result.LastInsertId())


	// 事物
	tx, err := database.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()	// 回滚操作，如果执行过程中没有错误，则需要在函数返回前执行tx.Commit()
				// 因此已经提交，所以回滚操作并不影响

	stmt, _ := tx.Prepare("your sql")		// 如果执行批量操作，则建议先把sql通过prepare()处理

	// 多行查询
	rows, _ := stmt.Query("value1", "value2")	// 这里不用再次加载sql语句，因为prepare()中已经处理过
	defer rows.Close()

	for rows.Next() {			// 遍历查询结果
		var value1, value2 string
		rows.Scan(&value1, &value2)	// 将查询结果赋给对应的变量
	}
	tx.Commit()
	// 事物结束
}
