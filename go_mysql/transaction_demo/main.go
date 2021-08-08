package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:12345678@tcp(127.0.0.1)/go_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func transactionDemo()  {
	tx, err := db.Begin()	// 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}

	sqlStr1 := `update user set age = age - 2 where id = 1`
	result1, err := tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}

	rowsAffected1, err := result1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := `update user set age = age + 2 where id = 2`
	result2, err := tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}

	rowsAffected2, err := result2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret2.RowsAffected() failed, err:%v\n", err)
		return
	}

	if rowsAffected1 == 1 && rowsAffected2 == 1 {
		 tx.Commit()
		 fmt.Printf("事务已提交\n")
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦")
	}
	fmt.Println("事务执行成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("连接数据库失败，err:%v\n", err)
	}
	transactionDemo()

}
