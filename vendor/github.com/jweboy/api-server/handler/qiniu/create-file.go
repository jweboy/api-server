package qiniu

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/handler"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/qiniu/api.v7/storage"
)

// UploadFile 文件上传
// @Summary 文件上传
// @Description 支持任何格式的文件上传，文件大小有限定
// @Tags qiniu
// @Accept  json
// @Produce  json
// @Param   bucketName     path    string     true        "存储空间名称"
// @Router /qiniu/file/{bucketName} [post]
func UploadFile(c *gin.Context) {
	// TODO: 文件大小需要作限制
	// TODO: 请求参数校验整理

	bucket := c.Param("bucket")
	if bucket == "" {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 获取表单提交的文件
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Print(err.Error())
		SendResponse(c, errno.ErrFileUpload, nil)
		return
	}

	// 读取文件的具体内容
	srcFile, err := file.Open()
	if err != nil {
		SendResponse(c, errno.ErrFileUpload, nil)
		return
	}

	// 文件名
	var fileName = file.Filename

	// 在files目录下生成新文件
	outFile, err := os.Create("files/" + fileName)
	if err != nil {
		SendResponse(c, errno.ErrFileUpload, nil)
		return
	}
	defer outFile.Close()

	// 拷贝源文件内容到新文件
	io.Copy(outFile, srcFile)

	// 设置基础配置
	cfg := getCfg()

	// 用于存储上传成功后的返回数据
	putRet := storage.PutRet{}

	// 文件上传需要增加的一些额外选项
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": fileName,
		},
	}

	// 获取上传的token
	// TODO: 七牛云有一个getToken的方法可以替换
	upToken := getToken(bucket)

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)

	// 上传文件到七牛云
	// FIXME: 由于目前对于文件流操作不熟悉，暂时采用上传到服务器之后再传一份到七牛云服务器，后期优化为数据流上传的方式。
	errs := formUploader.PutFile(context.Background(), &putRet, upToken, fileName, "files/"+fileName, &putExtra)
	if errs != nil {
		fmt.Println("errs:", errs.Error())
		SendResponse(c, errno.ErrFileUpload, nil)
		return
	}

	// 存入数据库的数据模型
	f := model.FileModel{
		Name: url.QueryEscape(putRet.Key),
		Key:  putRet.Hash,
	}

	// TODO: 数据库重复存入的问题
	// if err := f.Find(); err != nil {
	// 	fmt.Printf("Database find error => %v", err)
	// }

	// 存入数据库
	if errs := f.Create(); errs != nil {
		fmt.Printf("Database create error => %v", errs)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// 返回成功结果
	SendResponse(c, nil, nil)
}

// ListFile 获取指定空间的文件列表
func ListFile(c *gin.Context) {
	bucketName := c.Query("bucketName")
	size := c.Query("size")
	page := c.Query("page")
	if bucketName == "" || size == "" || page == "" {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// TODO: 直接从数据库获取可以进行分页操作，不从七牛云获取
}
