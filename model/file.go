package model

type FileModel struct {
	FileBaseModel
	Name string `json:"name" grom:"column: name;not null" binding:"required"`
	Key  string `json:"key" grom:"column: key;not null" binding:"required"`
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

func DeleteFile(id uint64) error {
	file := FileModel{}
	file.FileBaseModel.Id = id

	return DB.Self.Delete(&file).Error
}
