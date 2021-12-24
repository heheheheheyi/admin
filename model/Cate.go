package model

import (
	"admin/utils/errmsg"
	"database/sql"
	"fmt"
)

//资源数组
type Cate struct {
	Id       int    `json:"id"`
	Catename string `json:"catename"`
	Num      int    `json:"num"`
}

//查询资源名是否存在
func CheckCate(name string) int {
	var sqlname string
	sqlstr := "select catename from cate where catename=? "
	_ = db.QueryRow(sqlstr, name).Scan(&sqlname)
	if sqlname != "" {
		return errmsg.ErrorCatenameUsed
	}
	return errmsg.SUCCESS
}

//编辑时查询资源名是否存在（是否是自己或者修改的用户名已存在）
func EditCheckCate(id int, name string) int {
	var sqlid int
	sqlstr := "select id from cate where catename=? "
	_ = db.QueryRow(sqlstr, name).Scan(&sqlid)
	if sqlid > 0 && sqlid != id {
		return errmsg.ErrorCatenameUsed
	}
	return errmsg.SUCCESS
}

//新增分类
func CreateCate(data *Cate) int {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	if CheckCate(data.Catename) == errmsg.ErrorCatenameUsed {
		_ = tx.Rollback()
		return errmsg.ErrorCatenameUsed
	}
	sqlstr := "insert into cate(catename,num) values(?,0)"
	_, err = db.Exec(sqlstr, data.Catename)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//查询单个资源
func GetOneCate(id int) (Cate, int) {
	var sqlstr string
	var cate Cate
	sqlstr = "SELECT * FROM cate where id=?"
	_ = db.QueryRow(sqlstr, id).Scan(
		&cate.Id,
		&cate.Catename,
		&cate.Num,
	)
	return cate, errmsg.SUCCESS
}

//查询分类列表
func GetCate(catename string, pageSize int, pageNum int) ([]Cate, int, int) {
	var sqlstr string
	var cates []Cate
	var rows *sql.Rows
	var total int
	if catename == "" {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM cate"
		} else {
			sqlstr = fmt.Sprintf(
				"SELECT * FROM cate order by id limit %d,%d",
				(pageNum-1)*pageSize,
				pageSize,
			)
		}
		rows, err = db.Query(sqlstr)
		if err != nil {
			return cates, 0, errmsg.ErrorSql
		}
		sqlstr = "SELECT count(id) FROM cate"
		_ = db.QueryRow(sqlstr).Scan(&total)
	} else {
		if pageSize == 0 {
			sqlstr = fmt.Sprintf(
				"SELECT * FROM cate where catename like %s",
				"'%"+catename+"%'",
			)
		} else {
			sqlstr = fmt.Sprintf(
				"SELECT * FROM cate where catename like %s order by id limit %d,%d",
				"'%"+catename+"%'",
				(pageNum-1)*pageSize,
				pageSize,
			)
		}
		rows, err = db.Query(sqlstr)
		if err != nil {
			return cates, 0, errmsg.ErrorSql
		}
		sqlstr = fmt.Sprintf("SELECT count(id) FROM cate where catename like %s", "'%"+catename+"%'")
		_ = db.QueryRow(sqlstr).Scan(&total)
	}
	defer rows.Close()
	for rows.Next() {
		var cate Cate
		_ = rows.Scan(
			&cate.Id,
			&cate.Catename,
			&cate.Num,
		)
		cates = append(cates, cate)
	}
	return cates, total, errmsg.SUCCESS
}

//编辑分类
func EditCate(id int, data *Cate) int {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	code := EditCheckCate(id, data.Catename)
	if code == errmsg.ErrorCatenameUsed {
		_ = tx.Rollback()
		return code
	}
	if data.Catename == "" {
		_ = tx.Rollback()
		return errmsg.ERROR
	}
	sqlstr := "update cate set catename=? where id=?"
	_, err = db.Exec(sqlstr, data.Catename, id)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//删除分类
func DeleteCate(id int) int {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	cate, _ := GetOneCate(id)
	if cate.Num != 0 {
		_ = tx.Rollback()
		return errmsg.ErrorCateHasRes
	}
	sqlstr := "delete from cate where id=?"
	_, err = db.Exec(sqlstr, id)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}
