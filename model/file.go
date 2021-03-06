package model

import (
	"sync"
	"time"

	"github.com/jweboy/api-server/pkg/setting"
)

type UpdateDetailModel struct {
	FileName string `json:"name" example:"xx.png"`
	ID       int    `json:"id" example:"1"`
}

type FileModel struct {
	Id        uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	CreatedAt time.Time `json:"create_time" gorm:"column:createdAt"`
	Name      string    `json:"name" grom:"column: name;not null" binding:"required"`
	Key       string    `json:"key" grom:"column: key;not null" binding:"required"`
	Bucket    string    `json:"bucket" grom:"column: bucket;not null" binding:"required"`
	Size      int       `json:"size" grom:"column: size;not null" binding:"required"`
	MimeType  string    `json:"mime_type" grom:"column: mimeType;not null" binding:"required"`
}

type FileList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*FileModel
}

func (f *FileModel) TableName() string {
	return "tb_files"
}

func (f *FileModel) Create() error {
	return DB.Self.Create(&f).Error
}

func (f *FileModel) Find() error {
	return DB.Self.Find(&f).Error
}

func (f *FileModel) Update() error {
	return DB.Self.Model(&f).Where("id = ?", f.Id).Updates(f).Error
}

func DeleteFile(id uint64) error {
	file := FileModel{}
	file.Id = id

	return DB.Self.Delete(&file).Error
}

// ListFile 获取文件列表（带分页）
func ListFile(bucket string, offset, limit int) ([]*FileModel, uint64, error) {
	// TODO: 这里可能需要根据前面判断调整做调整
	if limit == 0 {
		limit = setting.AppSetting.PageSize
	}

	files := make([]*FileModel, 0)
	var count uint64

	// 第一次sql查询列表总数量
	if err := DB.Self.Model(&FileModel{}).Where("bucket=?", bucket).Count(&count).Error; err != nil {
		return files, count, err
	}

	// 第二次sql查询具体分页数据
	if err := DB.Self.Where("bucket=?", bucket).Offset((offset - 1) * limit).Limit(limit).Find(&files).Order("id desc").Error; err != nil {
		return files, count, err
	}

	// 返回分页数据和列表总数量
	return files, count, nil
}

// FileDetail 获取文件详情
func FileDetail(id int) (*FileModel, error) {
	file := &FileModel{}
	result := DB.Self.Where("id=?", id).First(&file)
	return file, result.Error
}
