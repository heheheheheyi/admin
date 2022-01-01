package model

import (
	"admin/utils/errmsg"
	"database/sql"
)

//资源数组
type Record struct {
	Id          int    `json:"id"`
	Rid         int    `json:"rid"`
	Cid         int    `json:"cid"`
	Uid         int    `json:"uid"`
	Mid         int    `json:"mid"`
	Time        string `json:"time"`
	ResInfo     Res
	UserInfo    User
	ManagerInfo User
	CateInfo    Cate
}

//查询借出记录
func GetBorrowRecord(pageSize int, pageNum int) ([]Record, int, int) {
	var sqlstr string
	var records []Record
	var rows *sql.Rows
	var total int
	if pageSize == 0 {
		sqlstr = "SELECT * FROM borrow_record order by id desc"
		rows, err = db.Query(sqlstr)
		if err != nil {
			return records, 0, errmsg.ErrorSql
		}
	} else {
		sqlstr = "SELECT * FROM borrow_record order by id desc limit ?,?"
		rows, err = db.Query(sqlstr, (pageNum-1)*pageSize, pageSize)
		if err != nil {
			return records, 0, errmsg.ErrorSql
		}
	}
	sqlstr = "SELECT count(id) FROM borrow_record"
	_ = db.QueryRow(sqlstr).Scan(&total)
	defer rows.Close()
	for rows.Next() {
		var record Record
		_ = rows.Scan(
			&record.Id,
			&record.Rid,
			&record.Cid,
			&record.Uid,
			&record.Mid,
			&record.Time,
		)
		var code int
		record.ResInfo, code = GetOneRes(record.Rid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		record.UserInfo, code = GetOneUser(record.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		record.ManagerInfo, code = GetOneUser(record.Mid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		record.CateInfo, _ = GetOneCate(record.Cid)
		records = append(records, record)
	}
	return records, total, errmsg.SUCCESS
}

//查询归还记录
func GetReturnRecord(pageSize int, pageNum int) ([]Record, int, int) {
	var sqlstr string
	var records []Record
	var rows *sql.Rows
	var total int
	if pageSize == 0 {
		sqlstr = "SELECT * FROM return_record order by id desc"
		rows, err = db.Query(sqlstr)
		if err != nil {
			return records, 0, errmsg.ErrorSql
		}
	} else {
		sqlstr = "SELECT * FROM return_record order by id desc limit ?,?"
		rows, err = db.Query(sqlstr, (pageNum-1)*pageSize, pageSize)
		if err != nil {
			return records, 0, errmsg.ErrorSql
		}
	}
	sqlstr = "SELECT count(id) FROM return_record"
	_ = db.QueryRow(sqlstr).Scan(&total)
	defer rows.Close()
	for rows.Next() {
		var record Record
		_ = rows.Scan(
			&record.Id,
			&record.Rid,
			&record.Cid,
			&record.Uid,
			&record.Mid,
			&record.Time,
		)
		var code int
		record.ResInfo, code = GetOneRes(record.Rid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		record.UserInfo, code = GetOneUser(record.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		record.ManagerInfo, code = GetOneUser(record.Mid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		record.CateInfo, _ = GetOneCate(record.Cid)
		records = append(records, record)
	}
	return records, total, errmsg.SUCCESS
}
