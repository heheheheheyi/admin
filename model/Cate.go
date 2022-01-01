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
	_ = db.QueryRow("select catename from cate where catename=? ", name).Scan(&sqlname)
	if sqlname != "" {
		return errmsg.ErrorCatenameUsed
	}
	return errmsg.SUCCESS
}

//编辑时查询资源名是否存在（是否是自己或者修改的用户名已存在）
func EditCheckCate(id int, name string) int {
	var sqlid int
	_ = db.QueryRow("select id from cate where catename=? ", name).Scan(&sqlid)
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
	_, err = db.Exec("insert into cate(catename,num) values(?,0)", data.Catename)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//查询单个资源
func GetOneCate(id int) (Cate, int) {
	var cate Cate
	_ = db.QueryRow("SELECT * FROM cate where id=?", id).Scan(
		&cate.Id,
		&cate.Catename,
		&cate.Num,
	)
	return cate, errmsg.SUCCESS
}

//查询分类列表
func GetCate(catename string, pageSize int, pageNum int) ([]Cate, int, int) {
	var cates []Cate
	var rows *sql.Rows
	var total int
	if catename == "" {
		if pageSize == 0 {
			rows, err = db.Query("SELECT * FROM cate")
			if err != nil {
				fmt.Println(err)
				return cates, 0, errmsg.ErrorSql
			}
		} else {
			rows, err = db.Query("SELECT * FROM cate order by id limit ?,?", (pageNum-1)*pageSize, pageSize)
			if err != nil {
				fmt.Println(err)
				return cates, 0, errmsg.ErrorSql
			}
		}
		_ = db.QueryRow("SELECT count(id) FROM cate").Scan(&total)
	} else {
		if pageSize == 0 {
			rows, err = db.Query("SELECT * FROM cate where catename like ?", "%"+catename+"%")
			if err != nil {
				fmt.Println(err)
				return cates, 0, errmsg.ErrorSql
			}
		} else {
			rows, err = db.Query("SELECT * FROM cate where catename like ? order by id limit ?,?", "%"+catename+"%", (pageNum-1)*pageSize, pageSize)
			if err != nil {
				fmt.Println(err)
				return cates, 0, errmsg.ErrorSql
			}
		}
		_ = db.QueryRow("SELECT count(id) FROM cate where catename like ?", "%"+catename+"%").Scan(&total)
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
	_, err = db.Exec("update cate set catename=? where id=?", data.Catename, id)
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
	_, err = db.Exec("delete from cate where id=?", id)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}
