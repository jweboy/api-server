package qiniu

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/spf13/viper"
	// . "github.com/jweboy/api-server/handler"
)

func Upload(c *gin.Context) {
	cfg := getCfg()
	mac := qbox.NewMac(viper.GetString("qiniu.accessKey"), viper.GetString("qiniu.secretKey"))
	bucketManger := storage.NewBucketManager(mac, &cfg)

	// 设置镜像存储
	prefix, delimiter, marker := "", "", ""
	entries, err := bucketManger.ListBucket("test", prefix, delimiter, marker)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ListBucket: %v\n", err)
		os.Exit(1)
	}
	for listItem := range entries {
		fmt.Println(listItem.Marker)
		fmt.Println(listItem.Item)
		fmt.Println(listItem.Dir)
		fmt.Println(strings.Repeat("-", 30))
	}

	// file, err := c.FormFile("file")

	// // out, err := os.Create("./files/" + file.Filename)
	// if err != nil {
	// 	fmt.Println(err.Error(), file.Filename)
	// }
	// // defer out.Close()
	// buf1 := bytes.NewBufferString("hello")
	// fmt.Print(buf1)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	SendResponse(c, errno.ErrFileUpload, nil)
	// 	return
	// }

	// cfg := getCfg()

	// ret := storage.PutRet{}
	// putExtra := storage.PutExtra{
	// 	Params: map[string]string{
	// 		"x:name": "test picture",
	// 	},
	// }

	// data := []byte("hello, this is qiniu cloud")

	// // dataLen := int64(len(data))
	// upToken := getToken("test")
	// // log.Info(token)

	// // 构建表单上传的对象
	// formUploader := storage.NewFormUploader(&cfg)
	// //putExtra.NoCrc32Check = true
	// errs := formUploader.Put(context.Background(), &ret, upToken, "image/", bytes.NewReader(data), file.Size, &putExtra)
	// if errs != nil {
	// 	fmt.Println("errs:", errs.Error())
	// 	return
	// }
	// fmt.Println("\nok", ret.Key, ret.Hash)
	// // for _, file := range files {
	// // 	fmt.Print(file.Filename)
	// // }

	// SendResponse(c, nil, "upload is ok")

	// name := c.Param("name")
	// log.Info(name)
	// SendResponse(c, nil, name)
}
