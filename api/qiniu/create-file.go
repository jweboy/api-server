package qiniu

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/jweboy/api-server/util"
	"github.com/qiniu/api.v7/storage"
)

// CreateQuery 文件上传请求 query
type CreateQuery struct {
	Bucket string `form:"bucket" binding:"required"` // 必须绑定对应的 bucket
}

// PutRet 七牛云返回成功后的数据结构
type PutRet struct {
	Key      string
	Hash     string
	Fsize    int
	Bucket   string
	Name     string
	MimeType string
}

// UploadFile 文件上传
// @Summary 文件上传
// @Description 支持任何格式的文件上传
// @Tags qiniu
// @Accept  multipart/form-data
// @Produce  json
// @Param   bucket   query    string     true        "存储空间名称"
// @Param	file	formData	file	true	"选择文件"
// @Router /qiniu/file [post]
// @Success 200
func UploadFile(c *gin.Context) {
	// TODO: 文件大小需要作限制
	var createQuery CreateQuery

	// 检查 bucket 字段是否上传且不为空
	if c.ShouldBindQuery(&createQuery) != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 获取表单提交的文件
	file, err := c.FormFile("file")

	// 检查 file 文件是否上传
	if err != nil {
		// fmt.Print(err.Error())
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	/* ============= 在 ./files 目录下生成新文件 ============= */
	// TODO: 文件上传改为字节或者数据 https://gist.github.com/ZenGround0/49e4a1aa126736f966a1dfdcb84abdae
	var fileName = file.Filename
	var bucket = createQuery.Bucket
	var mimeType = file.Header["Content-Type"][0]

	// 读取文件的具体内容
	srcFile, err := file.Open()
	if err != nil {
		SendResponse(c, errno.ErrFileUpload, nil)
		return
	}

	outFile, err := os.Create("files/" + fileName)
	if err != nil {
		SendResponse(c, errno.ErrFileUpload, nil)
		return
	}
	defer outFile.Close()

	// 拷贝源文件内容到新文件
	io.Copy(outFile, srcFile)

	/* ============= 上传到七牛云存储库 ============= */
	// TODO: 由于目前对于文件流操作不熟悉，暂时采用上传到服务器之后再传一份到七牛云服务器，后期优化为数据流上传的方式。

	// 文件上传需要增加的一些额外选项
	putExtra := storage.PutExtra{}

	// 用于存储上传成功后的返回数据
	putRet := PutRet{}

	// 获取上传的token
	uploadToken := util.GetToken(bucket)

	// 构建表单上传的对象
	formUploader := util.GetFormUploader()

	repErr := formUploader.PutFile(context.Background(), &putRet, uploadToken, fileName, "files/"+fileName, &putExtra)
	if repErr != nil {
		fmt.Println("repErr:", repErr.Error())
		SendResponse(c, errno.ErrFileUpload, nil)
		return
	}

	/* ============= 数据入库 ============= */
	// 定义入库数据模型
	f := model.FileModel{
		Name:     url.QueryEscape(putRet.Key),
		Key:      putRet.Hash,
		Bucket:   bucket,
		Size:     putRet.Fsize,
		MimeType: mimeType,
	}

	// 存入数据库
	if err := f.Create(); err != nil {
		fmt.Printf("Database create error => %v", err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// 返回成功结果
	SendResponse(c, nil, nil)
}
