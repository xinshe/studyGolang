package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	age  int
	name string
}

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

func queryRow(id int)  {
	sqlStr := `select name, age from user where id=?`
	var u user
	db.QueryRow(sqlStr, id).Scan(&u.name, &u.age)
	fmt.Println(u.name, u.age)
	fmt.Println(u)
}

func queryMore(id int)  {
	sqlStr := `select * from user where id > ?`
	rows, err := db.Query(sqlStr, 1)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func insert(u *user)  {
	sqlStr := `insert into user(id, name, age) values(?, ?, ?)`
	result, err := db.Exec(sqlStr, u.id, u.name, u.age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", insertId)
}

func update(id, age int) {
	sqlStr := `update user set age=? where id=?`
	result, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func delete(id int) {
	sqlStr := `delete from user where id = ?`
	result, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func prepareQuery()  {
	sqlStr := `select id, name, age from user where id > ?`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(2)
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func prepareInsert()  {
	sqlStr := `insert into user(id, name, age) values(?, ?, ?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	m := map[int]*user {
		5: &user{
			id:   5,
			age:  23,
			name: "Jack",
		},
		6: &user{
			id:   6,
			age:  20,
			name: "Emrys",
		},
	}
	for _, v := range m {
		stmt.Exec(&v.id, &v.name, &v.age)
	}
}

func main()  {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	fmt.Println("数据库连接成功")
	//queryRow(1)
	//queryMore(1)

	//u := &user{
	//	id:   6,
	//	age:  26,
	//	name: "Sunny",
	//}
	//insert(u)

	//update(6, 30)

	//delete(6)

	//prepareQuery()

	//prepareInsert()
	
}
