package test

import (
	"fmt"
	"goskeleton/app/utils/sql_factory"
	_ "goskeleton/bootstrap"
	"testing"
)

//	测试 sqlserver 之前，首先请去  app/utils/sql_factory/client.go 第 6 行， 打开 被注释的驱动，否则 sqlserver 无法操作
// database/db_demo_sqlserver 有最简洁的创建表命令,您可以快速初始化一个 sqlserver 数据库 db_goskeleton
// 本次测试使用最快捷的方式，只要保证 sqlserver 驱动初始化 ok 以及连接有效即可
// 实际应用请在 app/model 里面建表，整个操作与 mysql 类似

// 查询类
func TestSelect(t *testing.T) {

	sqlservConn := sql_factory.GetOneSqlClient("sqlserver")
	if sqlservConn == nil {
		return
	}
	sql := "select   user_name,pass,sex,age,remark,created_at,updated_at  from tb_users "
	rows, err := sqlservConn.Query(sql)
	if err == nil {
		var userName, pass, sex, age, remark, createdAt, updatedAt string
		for rows.Next() {
			_ = rows.Scan(&userName, &pass, &sex, &age, &remark, &createdAt, &updatedAt)
			fmt.Println(userName, pass, sex, age, remark, createdAt, updatedAt)
		}
		_ = rows.Close()
	} else {
		fmt.Println(err.Error())
	}
}

//执行类： 以修改数据为例，其他类似
func TestUpdate(t *testing.T) {

	sqlservConn := sql_factory.GetOneSqlClient("sqlserver")
	if sqlservConn == nil {
		return
	}
	sql := "update   tb_users   set  created_at=getdate() ,updated_at=getdate() ,remark='数据修改测试'  where   id=3  "
	result, err := sqlservConn.Exec(sql)
	if err == nil {
		effectiveNum, err2 := result.RowsAffected()
		if err2 == nil {
			fmt.Println("修改数据音响行数：", effectiveNum)
		} else {
			t.Errorf("修改数据发生错误：%s", err2.Error())
		}

	} else {
		t.Errorf("执行sql发生错误：%s", err.Error())
	}
}
