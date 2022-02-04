package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//在项目中我们通常可能会使用database/sql连接MySQL数据库。sqlx可以认为是Go语言内置database/sql的超集，它在优秀的内置database/sql基础上提供了一组扩展。这些扩展中除了大家常用来查询的Get(dest interface{}, ...) error和Select(dest interface{}, ...) error外还有很多其他强大的功能

var db *sqlx.DB

//初始化数据库实例
func initDB() (err error) {
	dsn := "root:wangfei121@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	ID   int
	Age  int
	Name string
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 6)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	for index, u := range users {
		fmt.Printf("%d:%#v\n", index, u)
	}
}

/*
	sqlx中的exec方法与原生sql中的exec使用基本一致
*/

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

//NamedQuery 与DB.NamedExec同理，这里是支持查询
// 使用map做命名查询
func namedQueryByMap() {
	sqlStr := "SELECT * FROM user WHERE name=:name"
	// 使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "老康"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

// 使用结构体命名查询，根据结构体字段的 db tag进行映射
func namedQueryByStruct() {
	sqlStr := "SELECT * FROM user WHERE name=:name"
	u := user{
		Name: "老康",
	}
	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	rows, err := db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

//对于事务操作，我们可以使用sqlx中提供的db.Beginx()和tx.Exec()方法。示例代码如下
func transactionDemo() (err error) {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Printf("rollback reason:%v \n", err)
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "Update user set age=20 where id=?"

	rs, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "Update user set age=50 where id=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err != nil {
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return err
}

/*
bindvars（绑定变量）
	查询占位符?在内部称为bindvars（查询占位符）,它非常重要。你应该始终使用它们向数据库发送值，因为它们可以防止SQL注入攻击。database/sql不尝试对查询文本进行任何验证；它与编码的参数一起按原样发送到服务器。除非驱动程序实现一个特殊的接口，否则在执行之前，查询是在服务器上准备的。因此bindvars是特定于数据库的:

	MySQL中使用?
	PostgreSQL使用枚举的$1、$2等bindvar语法
	SQLite中?和$1的语法都支持
	Oracle中使用:name的语法

bindvars的一个常见误解是，它们用来在sql语句中插入值。它们其实仅用于参数化，不允许更改SQL语句的结构。例如，使用bindvars尝试参数化列或表名将不起作用：
	// ？不能用来插入表名（做SQL语句中表名的占位符）
	db.Query("SELECT * FROM ?", "mytable")

	// ？也不能用来插入列名（做SQL语句中列名的占位符）
	db.Query("SELECT ?, ? FROM people", "name", "location")
*/

//创建student表sql
/*
	CREATE TABLE `student` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
	)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
*/

//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
type student struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

//自行构建批量插入的语句的方法
func BatchInsertStudent(students []*student) error {
	sqlStrings := make([]string, 0, len(students))
	valueArgs := make([]interface{}, 0, len(students)*2)
	for _, s := range students {
		sqlStrings = append(sqlStrings, "(?, ?)")
		valueArgs = append(valueArgs, s.Name)
		valueArgs = append(valueArgs, s.Age)
	}
	sql := fmt.Sprintf("INSERT INTO student (name, age) VALUES %s", strings.Join(sqlStrings, ","))
	fmt.Println(sql)
	_, err := db.Exec(sql, valueArgs...)
	return err
}

//使用sqlx.In实现批量插入
//前提是需要我们的结构体实现driver.Valuer接口
func (s student) Value() (driver.Value, error) {
	return []interface{}{s.Name, s.Age}, nil
}

//使用sqlx.In实现批量插入代码如下
//BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertStudent2(students []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO student (name, age) VALUES (?), (?), (?)",
		students..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	_, err := db.Exec(query, args...)
	return err
}

//使用NamedExec实现批量插入
//注意 ：该功能需1.3.1版本以上，并且1.3.1版本目前还有点问题，sql语句最后不能有空格和;，详见issues/690。
//使用NamedExec实现批量插入的代码如下
// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertStudent3(students []*student) error {
	_, err := db.NamedExec("INSERT INTO student (name, age) VALUES (:name, :age)", students)
	return err
}

/*
sqlx.In的查询示例
	关于sqlx.In这里再补充一个用法，在sqlx查询语句中实现In查询和FIND_IN_SET函数。即实现SELECT * FROM user WHERE id in (3, 2, 1);和SELECT * FROM user WHERE id in (3, 2, 1) ORDER BY FIND_IN_SET(id, '3,2,1');。

in查询
*/

// QueryByIDs 根据给定ID查询
func QueryByIDs(ids []int) (students []student, err error) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT name, age FROM student WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	fmt.Println(query)
	fmt.Println(args)

	err = db.Select(&students, query, args...)
	return students, err
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int)(students []student, err error){
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT name, age FROM student WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}

	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&students, query, args...)
	return students, err
}

func main() {
	initDB()

	defer db.Close()
	// queryRowDemo()
	// queryMultiRowDemo()

	// insertRowDemo()
	// updateRowDemo()
	// deleteRowDemo()

	// namedQueryByMap()
	// namedQueryByStruct()
	// transactionDemo()

	// s1 := student{Name: "student_1", Age: 18}
	// s2 := student{Name: "student_2", Age: 28}
	// s3 := student{Name: "student_3", Age: 38}

	//批量创建方式一
	// students1 := []*student{
	// 	&s1, &s2, &s3,
	// }
	// err1 := BatchInsertStudent(students1)
	// if err1 != nil {
	// 	fmt.Println("BatchInsertStudent failed, err:%v\n", err1)
	// }

	//批量创建方法二
	// students2 := []interface{}{
	// 	s1, s2, s3,
	// }
	// err2 := BatchInsertStudent2(students2)
	// if err2 != nil {
	// 	fmt.Println("BatchInsertStudent2 failed, err:%v\n", err2)
	// }

	//批量创建方法三
	// students3 := []*student{
	// 	&s1, &s2, &s3,
	// }
	// err3 := BatchInsertStudent3(students3)
	// if err3 != nil {
	// 	fmt.Printf("BatchInsertStudent3 failed, err:%v\n", err3)
	// }

	// sqlx.In的查询示例
	// students, err := QueryByIDs([]int{31, 32, 33})
	// if err !=nil{
	// 	fmt.Printf("QueryByIDs err:%v",err)
	// }
	// for _, s := range students{
	// 	fmt.Printf("name:%v, age:%d \n",s.Name, s.Age)
	// }

	// in查询和FIND_IN_SET函数
	students, err := QueryAndOrderByIDs([]int{31, 32, 33})
	if err !=nil{
		fmt.Printf("QueryAndOrderByIDs err:%v",err)
	}
	for _, s := range students{
		fmt.Printf("name:%v, age:%d \n",s.Name, s.Age)
	}
}
