package aliyunoss

import (
	common "common/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

var EP string
var AK string
var SK string

func init()  {
	EP = common.ConfMap["EP"]
	AK = common.ConfMap["AK"]
	SK = common.ConfMap["SK"]
}

func UploadToOss(fileName string, path string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service failed: %s", err)
		return false
	}
	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bucket failed: %s", err)
		return false
	}
	err = bucket.UploadFile(fileName, path, 500*1024, oss.Routines(5))
	if err != nil {
		log.Printf("uploading object failed :%s", err)
		return false
	}
	return true
}

func DeleteObject(fileName string, bn string) bool {
	client, err := oss.New(EP,AK,SK)
	if err != nil {
		log.Printf("Init oss service error :%s ",err)
		return false
	}
	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bucket error:%s", err)
		return false
	}
	err = bucket.DeleteObject(fileName)
	if err != nil {
		log.Printf("delete object error:%s", err)
		return false
	}
	return true
}
