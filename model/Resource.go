package model

import (
	"admin/utils/errmsg"
	"database/sql"
	"fmt"
	"time"
)

//资源数组
type Res struct {
	Id         int    `json:"id"`
	Resname    string `json:"resname"`
	Cid        int    `json:"cid"`
	Img        string `json:"img"`
	Status     int    `json:"status"`
	Uid        int    `json:"uid"`
	BorrowTime string `json:"borrow_time"`
	ReturnTime string `json:"return_time"`
	UserInfo   User
	CateInfo   Cate
}

//查询资源名是否存在
func CheckRes(name string) int {
	var sqlname string
	sqlstr := "select resname from res where resname=? "
	_ = db.QueryRow(sqlstr, name).Scan(&sqlname)
	if sqlname != "" {
		return errmsg.ErrorResnameUsed
	}
	return errmsg.SUCCESS
}

//编辑时查询资源名是否存在（是否是自己或者修改的用户名已存在）
func EditCheckRes(id int, name string) int {
	var sqlid int
	sqlstr := "select id from res where resname=? "
	_ = db.QueryRow(sqlstr, name).Scan(&sqlid)
	if sqlid > 0 && sqlid != id {
		return errmsg.ErrorResnameUsed
	}
	return errmsg.SUCCESS
}

//新增资源
func CreateRes(data *Res) int {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	if CheckRes(data.Resname) == errmsg.ErrorResnameUsed {
		_ = tx.Rollback()
		return errmsg.ErrorResnameUsed
	}
	sqlstr := "insert into res(resname,cid,status,img) values(?,?,?,?)"
	_, err = db.Exec(sqlstr, data.Resname, data.Cid, 1, data.Img)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	sqlstr = "update cate set num=num+1 where id=?"
	_, err = db.Exec(sqlstr, data.Cid)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//查询单个资源
func GetOneRes(id int) (Res, int) {
	var sqlstr string
	var res Res
	sqlstr = "SELECT * FROM res where id=?"
	_ = db.QueryRow(sqlstr, id).Scan(
		&res.Id,
		&res.Resname,
		&res.Cid,
		&res.Img,
		&res.Status,
		&res.Uid,
		&res.BorrowTime,
		&res.ReturnTime,
	)
	var code int
	res.UserInfo, code = GetOneUser(res.Uid)
	if code != errmsg.SUCCESS {
		println(errmsg.GetErrMsg(code))
	}
	res.CateInfo, _ = GetOneCate(res.Cid)
	return res, errmsg.SUCCESS
}

//查询分类下的资源列表
func GetCateRes(cid int, resname string, pageSize int, pageNum int) ([]Res, int, int) {
	var sqlstr string
	var ress []Res
	var rows *sql.Rows
	var total int
	if resname == "" {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where cid = ?"
			rows, err = db.Query(sqlstr, cid)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where cid=? order by id limit ?,?"
			rows, err = db.Query(sqlstr, cid, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where cid=?"
		_ = db.QueryRow(sqlstr, cid).Scan(&total)
	} else {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where resname like ? and cid=?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", cid)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where resname like ? and cid=? order by id limit ?,?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", cid, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where resname like ? and cid=?"
		_ = db.QueryRow(sqlstr, "%"+resname+"%", cid).Scan(&total)
	}
	defer rows.Close()
	for rows.Next() {
		var res Res
		_ = rows.Scan(
			&res.Id,
			&res.Resname,
			&res.Cid,
			&res.Img,
			&res.Status,
			&res.Uid,
			&res.BorrowTime,
			&res.ReturnTime,
		)
		var code int
		res.UserInfo, code = GetOneUser(res.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		res.CateInfo, _ = GetOneCate(res.Cid)
		ress = append(ress, res)
	}
	return ress, total, errmsg.SUCCESS
}

//查询资源列表
func GetRes(resname string, pageSize int, pageNum int) ([]Res, int, int) {
	var sqlstr string
	var ress []Res
	var rows *sql.Rows
	var total int
	if resname == "" {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res"
			rows, err = db.Query(sqlstr)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res order by id limit ?,?"
			rows, err = db.Query(sqlstr, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}

		}
		sqlstr = "SELECT count(id) FROM res"
		_ = db.QueryRow(sqlstr).Scan(&total)
	} else {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where resname like ?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where resname like ? order by id limit ?,?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where resname like ?"
		_ = db.QueryRow(sqlstr, "%"+resname+"%").Scan(&total)
	}
	defer rows.Close()
	for rows.Next() {
		var res Res
		_ = rows.Scan(
			&res.Id,
			&res.Resname,
			&res.Cid,
			&res.Img,
			&res.Status,
			&res.Uid,
			&res.BorrowTime,
			&res.ReturnTime,
		)
		var code int
		res.UserInfo, code = GetOneUser(res.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		res.CateInfo, _ = GetOneCate(res.Cid)
		ress = append(ress, res)
	}
	return ress, total, errmsg.SUCCESS
}

//查询未借出的资源
func GetUnBorrow(resname string, pageSize int, pageNum int) ([]Res, int, int) {
	var sqlstr string
	var ress []Res
	var rows *sql.Rows
	var total int
	if resname == "" {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where status=1"
			rows, err = db.Query(sqlstr)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where status=1 order by id limit ?,?"
			rows, err = db.Query(sqlstr, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where status=1"
		_ = db.QueryRow(sqlstr).Scan(&total)
	} else {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where resname like ? and status=1"
			rows, err = db.Query(sqlstr, "%"+resname+"%")
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}

		} else {
			sqlstr = "SELECT * FROM res where resname like ? and status=1 order by id limit ?,?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where resname like ? and status=1"
		_ = db.QueryRow(sqlstr, "%"+resname+"%").Scan(&total)
	}
	defer rows.Close()
	for rows.Next() {
		var res Res
		_ = rows.Scan(
			&res.Id,
			&res.Resname,
			&res.Cid,
			&res.Img,
			&res.Status,
			&res.Uid,
			&res.BorrowTime,
			&res.ReturnTime,
		)
		var code int
		res.UserInfo, code = GetOneUser(res.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		res.CateInfo, _ = GetOneCate(res.Cid)
		ress = append(ress, res)
	}
	return ress, total, errmsg.SUCCESS
}

//查询待确认借出的资源
func GetApplyBorrow(resname string, pageSize int, pageNum int) ([]Res, int, int) {
	var sqlstr string
	var ress []Res
	var rows *sql.Rows
	var total int
	if resname == "" {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where status=2 and uid!=0"
			rows, err = db.Query(sqlstr)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where status=2 and uid!=0 order by id limit ?,?"
			rows, err = db.Query(sqlstr, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where status=2 and uid!=0"
		_ = db.QueryRow(sqlstr).Scan(&total)
	} else {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where resname like ? and status=2 and uid!=0"
			rows, err = db.Query(sqlstr, "%"+resname+"%")
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where resname like ? and status=2 and uid!=0 order by id limit ?,?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		rows, err = db.Query(sqlstr)
		if err != nil {
			return ress, 0, errmsg.ErrorSql
		}
		sqlstr = "SELECT count(id) FROM res where resname like ? and status=2 and uid!=0"
		_ = db.QueryRow(sqlstr, "%"+resname+"%").Scan(&total)
	}
	defer rows.Close()
	for rows.Next() {
		var res Res
		_ = rows.Scan(
			&res.Id,
			&res.Resname,
			&res.Cid,
			&res.Img,
			&res.Status,
			&res.Uid,
			&res.BorrowTime,
			&res.ReturnTime,
		)
		var code int
		res.UserInfo, code = GetOneUser(res.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		res.CateInfo, _ = GetOneCate(res.Cid)
		ress = append(ress, res)
	}
	return ress, total, errmsg.SUCCESS
}

//查询借出的资源
func GetBorrowRes(resname string, pageSize int, pageNum int) ([]Res, int, int) {
	var sqlstr string
	var ress []Res
	var rows *sql.Rows
	var total int
	if resname == "" {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where status=3 and uid!=0"
			rows, err = db.Query(sqlstr)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where status=3 and uid!=0 order by id limit ?,?"
			rows, err = db.Query(sqlstr, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where status=3 and uid!=0"
		_ = db.QueryRow(sqlstr).Scan(&total)
	} else {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where resname like ? and status=3 and uid!=0"
			rows, err = db.Query(sqlstr, "%"+resname+"%")
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where resname like ? and status=3 and uid!=0 order by id limit ?,?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		rows, err = db.Query(sqlstr)
		if err != nil {
			return ress, 0, errmsg.ErrorSql
		}
		sqlstr = "SELECT count(id) FROM res where resname like ? and status=3 and uid!=0"
		_ = db.QueryRow(sqlstr, "%"+resname+"%").Scan(&total)
	}
	defer rows.Close()
	for rows.Next() {
		var res Res
		_ = rows.Scan(
			&res.Id,
			&res.Resname,
			&res.Cid,
			&res.Img,
			&res.Status,
			&res.Uid,
			&res.BorrowTime,
			&res.ReturnTime,
		)
		var code int
		res.UserInfo, code = GetOneUser(res.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		res.CateInfo, _ = GetOneCate(res.Cid)
		ress = append(ress, res)
	}
	return ress, total, errmsg.SUCCESS
}

//查询未借出的资源（分类）
func GetCateUnBorrow(cid int, resname string, pageSize int, pageNum int) ([]Res, int, int) {
	var sqlstr string
	var ress []Res
	var rows *sql.Rows
	var total int
	if resname == "" {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where status=1 and cid=?"
			rows, err = db.Query(sqlstr, cid)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where status=1 and cid=? order by id limit ?,?"
			rows, err = db.Query(sqlstr, cid, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where status=1 and cid=?"
		_ = db.QueryRow(sqlstr, cid).Scan(&total)
	} else {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where resname like ? and status=1 and cid=?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", cid)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where resname like ? and status=1 and cid=? order by id limit ?,?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", cid, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where resname like ? and status=1 and cid=?"
		_ = db.QueryRow(sqlstr, "%"+resname+"%", cid).Scan(&total)
	}
	defer rows.Close()
	for rows.Next() {
		var res Res
		_ = rows.Scan(
			&res.Id,
			&res.Resname,
			&res.Cid,
			&res.Img,
			&res.Status,
			&res.Uid,
			&res.BorrowTime,
			&res.ReturnTime,
		)
		var code int
		res.UserInfo, code = GetOneUser(res.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		res.CateInfo, _ = GetOneCate(res.Cid)
		ress = append(ress, res)
	}
	return ress, total, errmsg.SUCCESS
}

//查询待确认借出的资源（分类）
func GetCateApplyBorrow(cid int, resname string, pageSize int, pageNum int) ([]Res, int, int) {
	var sqlstr string
	var ress []Res
	var rows *sql.Rows
	var total int
	if resname == "" {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where status=2 and uid!=0 and cid=?"
			rows, err = db.Query(sqlstr, cid)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where status=2 and uid!=0 and cid=? order by id limit ?,?"
			rows, err = db.Query(sqlstr, cid, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where status=2 and uid!=0 and cid=?"
		_ = db.QueryRow(sqlstr, cid).Scan(&total)
	} else {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where resname like ? and status=2 and uid!=0 and cid=?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", cid)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where resname like ? and status=2 and uid!=0 and cid=? order by id limit ?,?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", cid, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where resname like ? and status=2 and uid!=0 and cid=?"
		_ = db.QueryRow(sqlstr, "%"+resname+"%", cid).Scan(&total)
	}
	defer rows.Close()
	for rows.Next() {
		var res Res
		_ = rows.Scan(
			&res.Id,
			&res.Resname,
			&res.Cid,
			&res.Img,
			&res.Status,
			&res.Uid,
			&res.BorrowTime,
			&res.ReturnTime,
		)
		var code int
		res.UserInfo, code = GetOneUser(res.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		res.CateInfo, _ = GetOneCate(res.Cid)
		ress = append(ress, res)
	}
	return ress, total, errmsg.SUCCESS
}

//查询借出的资源（分类）
func GetCateBorrowRes(cid int, resname string, pageSize int, pageNum int) ([]Res, int, int) {
	var sqlstr string
	var ress []Res
	var rows *sql.Rows
	var total int
	if resname == "" {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where status=3 and uid!=0 and cid=?"
			rows, err = db.Query(sqlstr, cid)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where status=3 and uid!=0 and cid=? order by id limit ?,?"
			rows, err = db.Query(sqlstr, cid, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where status=3 and uid!=0 and cid=?"
		_ = db.QueryRow(sqlstr, cid).Scan(&total)
	} else {
		if pageSize == 0 {
			sqlstr = "SELECT * FROM res where resname like ? and status=3 and uid!=0 and cid=?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", cid)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		} else {
			sqlstr = "SELECT * FROM res where resname like ? and status=3 and uid!=0 and cid=? order by id limit ?,?"
			rows, err = db.Query(sqlstr, "%"+resname+"%", cid, (pageNum-1)*pageSize, pageSize)
			if err != nil {
				return ress, 0, errmsg.ErrorSql
			}
		}
		sqlstr = "SELECT count(id) FROM res where resname like ? and status=3 and uid!=0 and cid=?"
		_ = db.QueryRow(sqlstr, "%"+resname+"%", cid).Scan(&total)
	}
	defer rows.Close()
	for rows.Next() {
		var res Res
		_ = rows.Scan(
			&res.Id,
			&res.Resname,
			&res.Cid,
			&res.Img,
			&res.Status,
			&res.Uid,
			&res.BorrowTime,
			&res.ReturnTime,
		)
		var code int
		res.UserInfo, code = GetOneUser(res.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		res.CateInfo, _ = GetOneCate(res.Cid)
		ress = append(ress, res)
	}
	return ress, total, errmsg.SUCCESS
}

//编辑资源
func EditRes(id int, data *Res) int {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	code := EditCheckRes(id, data.Resname)
	if code == errmsg.ErrorResnameUsed {
		_ = tx.Rollback()
		return code
	}
	if data.Resname == "" || data.Cid == 0 {
		_ = tx.Rollback()
		return errmsg.ERROR
	}
	sqlstr := "update res set resname=?,cid=? where id=?"
	_, err = db.Exec(sqlstr, data.Resname, data.Cid, id)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//删除资源
func DeleteRes(id int) int {
	var cid int
	var status int
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	sqlstr := "select status,cid from res where id=?"
	err = db.QueryRow(sqlstr, id).Scan(&status, &cid)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	if status != 1 {
		return errmsg.ErrorResHasBorrow
	}
	sqlstr = "delete from res where id=?"
	_, err = db.Exec(sqlstr, id)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	sqlstr = "update cate set num=num-1 where id=?"
	_, err = db.Exec(sqlstr, cid)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//确认借出资源
func BorrowRes(id int, mid int) int {
	var res Res
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	sqlstr := "select * from res where id=?"
	err = db.QueryRow(sqlstr, id).Scan(
		&res.Id,
		&res.Resname,
		&res.Cid,
		&res.Img,
		&res.Status,
		&res.Uid,
		&res.BorrowTime,
		&res.ReturnTime,
	)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	sqlstr = "update res set status=? where id=?"
	_, err = db.Exec(sqlstr, 3, id)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	sqlstr = "update users set num=num+1 where id=?"
	_, err = db.Exec(sqlstr, res.Uid)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	sqlstr = "insert into borrow_record(rid, cid, uid, mid, time) value (?,?,?,?,?)"
	_, err = db.Exec(sqlstr, id, res.Cid, res.Uid, mid, time.Now().In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05"))
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//拒绝借出资源
func RejectBorrow(id int, mid int) int {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	sqlstr := "update res set status=?,uid=null,borrow_time=null,return_time=null where id=?"
	_, err = db.Exec(sqlstr, 1, id)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//申请借出资源
func ApplyBorrowRes(id int, uid int, BorrowTime string, ReturnTime string) int {
	var status int
	var userid int
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	sqlstr := "SELECT status,uid FROM res where id=?"
	_ = db.QueryRow(sqlstr, id).Scan(&status, &userid)
	if status != 1 || userid > 0 {
		_ = tx.Rollback()
		return errmsg.ErrorResHasBorrow
	}
	sqlstr = "update res set uid=?,status=?,borrow_time=?,return_time=? where id=?"
	_, err = db.Exec(sqlstr, uid, 2, BorrowTime, ReturnTime, id)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//归还资源
func ReturnRes(id int, mid int) int {
	var res Res
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	sqlstr := "select * from res where id=?"
	err = db.QueryRow(sqlstr, id).Scan(
		&res.Id,
		&res.Resname,
		&res.Cid,
		&res.Img,
		&res.Status,
		&res.Uid,
		&res.BorrowTime,
		&res.ReturnTime,
	)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	sqlstr = "update res set status=1,uid=null,borrow_time=null,return_time=null where id=?"
	_, err = db.Exec(sqlstr, id)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	sqlstr = "update users set num=num-1 where id=?"
	_, err = db.Exec(sqlstr, res.Uid)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	sqlstr = "insert into return_record(rid, cid, uid, mid, time) value (?,?,?,?,?)"
	_, err = db.Exec(sqlstr, id, res.Cid, res.Uid, mid, time.Now().In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05"))
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}
