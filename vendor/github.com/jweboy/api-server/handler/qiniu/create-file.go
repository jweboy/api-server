package qiniu

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/handler"
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
	// TODO: 文件代销需要作限制

	// 检查对应的存储空间名是否上传
	bucketName := c.Param("bucketName")
	if bucketName == "" {
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
	upToken := getToken(bucketName)

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)

	// 上传文件
	// FIXME: 由于目前对于文件流操作不熟悉，暂时采用上传到服务器之后再传一份到七牛云服务器，后期优化为数据流上传的方式。
	errs := formUploader.PutFile(context.Background(), &putRet, upToken, "image/"+fileName, "files/"+fileName, &putExtra)
	if errs != nil {
		fmt.Println("errs:", errs.Error())
		SendResponse(c, errno.ErrFileUpload, nil)
		return
	}

	// TODO: 考虑增加一步数据库的入库操作

	// 返回成功结果
	SendResponse(c, nil, putRet)

}
