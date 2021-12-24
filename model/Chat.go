package model

import (
	"admin/utils/errmsg"
	"fmt"
)

//用户数组
type Chat struct {
	Id           int    `json:"id"`
	Uid          int    `json:"uid"`
	Fid          int    `json:"fid"`
	Content      string `json:"content"`
	Time         string `json:"time"`
	UserInfo     User
	ReceiverInfo User
}

//新增聊天
func AddChat(data *Chat) int {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return errmsg.ErrorSql
	}
	sqlstr := "insert into chat(uid, fid, content, time,isread) values(?,?,?,?,1)"
	_, err = db.Exec(sqlstr, data.Uid, data.Fid, data.Content, data.Time)
	if err != nil {
		_ = tx.Rollback()
		return errmsg.ErrorSql
	}
	_ = tx.Commit()
	return errmsg.SUCCESS
}

//查询未读聊天记录
func GetNoreadChat(id int) ([]Chat, int) {
	var sqlstr string
	var chats []Chat
	sqlstr = "SELECT id,uid,fid,content,time FROM chat where fid=? and isread=1"
	rows, err := db.Query(sqlstr, id)
	if err != nil {
		return nil, errmsg.ErrorSql
	}
	defer rows.Close()
	for rows.Next() {
		var chat Chat
		_ = rows.Scan(&chat.Id, &chat.Uid, &chat.Fid, &chat.Content, &chat.Time)
		var code int
		chat.UserInfo, code = GetOneUser(chat.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		chat.ReceiverInfo, code = GetOneUser(chat.Fid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		chats = append(chats, chat)
	}
	sqlstr = "update chat set isread=2 where fid=? "
	_, err = db.Exec(sqlstr, id)
	if err != nil {
		return nil, errmsg.ErrorSql
	}
	return chats, errmsg.SUCCESS
}

//查询聊天记录
func GetChat(id int) ([]Chat, int) {
	var sqlstr string
	var chats []Chat
	sqlstr = "SELECT id,uid,fid,content,time FROM chat where fid=? and isread=2 order by id desc "
	rows, err := db.Query(sqlstr, id)
	if err != nil {
		return nil, errmsg.ErrorSql
	}
	defer rows.Close()
	for rows.Next() {
		var chat Chat
		_ = rows.Scan(&chat.Id, &chat.Uid, &chat.Fid, &chat.Content, &chat.Time)
		var code int
		chat.UserInfo, code = GetOneUser(chat.Uid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		chat.ReceiverInfo, code = GetOneUser(chat.Fid)
		if code != errmsg.SUCCESS {
			println(errmsg.GetErrMsg(code))
		}
		chats = append(chats, chat)
	}
	return chats, errmsg.SUCCESS
}

//查询未读记录
func GetCount(id int) (int, int) {
	var sqlstr string
	var count int
	sqlstr = "select count(id) from chat where isread=1 and fid=? order by id desc "
	err = db.QueryRow(sqlstr, id).Scan(&count)
	if err != nil {
		return 0, errmsg.ErrorSql
	}
	return count, errmsg.SUCCESS
}
