package model

import (
	"admin/utils/errmsg"
	"database/sql"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"log"
)

//用户数组
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Num      int    `json:"num"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `json:"role" validate:"required,gte=1" label:"权限"`
}

//查询用户是否存在
func CheckUser(name string) int {
	sqlstr := "select id from users where username=? "
	row := db.QueryRow(sqlstr, name)
	var sqlid int
	_ = row.Scan(&sqlid)
	println(sqlid)
	if sqlid > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

//编辑时查询用户是否存在（是否是自己或者修改的用户名已存在）
func EditCheckUser(id int, name string) (code int) {
	sqlstr := "select id from users where username=? "
	row := db.QueryRow(sqlstr, name)
	var sqlid int
	_ = row.Scan(&sqlid)
	if sqlid > 0 && sqlid != id {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

//新增用户
func AddUser(data *User) int {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	if CheckUser(data.Username) == errmsg.ErrorUsernameUsed {
		_ = tx.Rollback()
		return errmsg.ErrorUsernameUsed
	}
	data.Password = ScryptPw(data.Password)
	sqlstr := "insert into users(username,num,password,role) values(?,?,?,?)"
	_, err = db.Exec(sqlstr, data.Username, 0, data.Password, data.Role)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//查询用户id
func GetUserId(username string) (int, int) {
	var sqlstr string
	var id int
	sqlstr = "SELECT id FROM users where username=?"
	row := db.QueryRow(sqlstr, username)
	_ = row.Scan(&id)
	return id, errmsg.SUCCESS
}

//查询单个用户
func GetOneUser(id int) (User, int) {
	var sqlstr string
	var user User
	sqlstr = "SELECT id, username,num, password, role FROM users where id=?"
	row := db.QueryRow(sqlstr, id)
	_ = row.Scan(&user.Id, &user.Username, &user.Num, &user.Password, &user.Role)
	return user, errmsg.SUCCESS
}

//查询用户列表
func GetUser(username string, pageSize int, pageNum int) ([]User, int, int) {
	var sqlstr string
	var users []User
	var rows *sql.Rows
	var total int
	if username == "" {
		if pageSize == 0 {
			sqlstr = "SELECT id, username,num, password, role FROM users"
		} else {
			sqlstr = fmt.Sprintf(
				"SELECT id, username,num, password, role FROM users order by id limit %d,%d",
				(pageNum-1)*pageSize,
				pageSize,
			)
		}
		rows, err = db.Query(sqlstr)
		if err != nil {
			return users, 0, errmsg.ErrorSql
		}
		sqlstr = "SELECT count(id) FROM users"
		_ = db.QueryRow(sqlstr).Scan(&total)
	} else {
		if pageSize == 0 {
			sqlstr = fmt.Sprintf(
				"SELECT id, username,num, password, role FROM users where username like %s",
				"'%"+username+"%'",
			)
		} else {
			sqlstr = fmt.Sprintf(
				"SELECT id,username,password,role FROM users where username like %s order by id limit %d,%d",
				"'%"+username+"%'",
				(pageNum-1)*pageSize,
				pageSize,
			)
		}
		rows, err = db.Query(sqlstr)
		if err != nil {
			return users, 0, errmsg.ErrorSql
		}
		sqlstr = fmt.Sprintf("SELECT count(id) FROM users where username like %s", "'%"+username+"%'")
		_ = db.QueryRow(sqlstr).Scan(&total)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		_ = rows.Scan(&user.Id, &user.Username, &user.Num, &user.Password, &user.Role)
		users = append(users, user)
	}
	return users, total, errmsg.SUCCESS
}

//编辑用户
func EditUser(id int, data *User) int {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	if EditCheckUser(id, data.Username) == errmsg.ErrorUsernameUsed {
		_ = tx.Rollback()
		return errmsg.ErrorUsernameUsed
	}
	if data.Role == 0 {
		sqlstr := "update users set username=? where id=?"
		_, err := db.Exec(sqlstr, data.Username, id)
		if err != nil {
			_ = tx.Rollback()
			return errmsg.ErrorSql
		}
	} else {
		sqlstr := "update users set username=?,role=? where id=?"
		_, err := db.Exec(sqlstr, data.Username, data.Role, id)
		if err != nil {
			_ = tx.Rollback()
			return errmsg.ErrorSql
		}
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//删除用户
func DeleteUser(id int) int {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	user, _ := GetOneUser(id)
	if user.Num != 0 {
		_ = tx.Rollback()
		return errmsg.ErrorUserHasBollow
	}
	sqlstr := "delete from users where id=?"
	_, err = db.Exec(sqlstr, id)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{11, 22, 33, 44, 55, 66, 77, 88}
	HashPw, err := scrypt.Key([]byte(password), salt, 16, 8, 1, KeyLen)
	if err != nil {
		log.Fatalln(err)
	}
	FinalPw := base64.StdEncoding.EncodeToString(HashPw)
	return FinalPw
}

//登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	code := CheckUser(username)
	if code == errmsg.SUCCESS {
		return user, errmsg.ErrorUserNotExist
	}
	sqlstr := "select id,username,password,role,num from users where username=? "
	row := db.QueryRow(sqlstr, username)
	_ = row.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.Num)
	if user.Password != ScryptPw(password) {
		return user, errmsg.ErrorPasswordWrong
	}
	if user.Role != 1 {
		return user, errmsg.ErrorUserNoRole
	}
	return user, errmsg.SUCCESS
}
