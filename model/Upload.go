package model

import (
	"admin/utils"
	"admin/utils/errmsg"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var Secretkey = utils.SecretKey
var Bucket = utils.Bucket
var Img = utils.YunServer

func UpLoadFile(file multipart.File, filesize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, Secretkey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(
		context.Background(),
		&ret,
		upToken,
		file,
		filesize,
		&putExtra,
	)
	if err != nil {
		return "", errmsg.ErrorUpload
	}
	url := Img + ret.Key
	return url, errmsg.SUCCESS
}
