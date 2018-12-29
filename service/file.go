package service

import (
	"sync"

	"github.com/jweboy/api-server/util"

	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
)

// TODO: 具体并发的过程理解加注释

func ListFile(bucket string, page, size int) ([]*model.FileModel, uint64, error) {
	list := make([]*model.FileModel, 0)

	files, count, err := model.ListFile(bucket, page, size)
	if err != nil {
		return nil, count, errno.ErrDatabase
	}

	ids := []uint64{}
	for _, file := range files {
		ids = append(ids, file.Id)
	}

	wg := sync.WaitGroup{}
	fileList := model.FileList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.FileModel, len(files)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, f := range files {
		wg.Add(1)

		go func(f *model.FileModel) {
			defer wg.Done()

			decodeName, err := util.DecodeStr(f.Name)
			if err != nil {
				errChan <- err
				return
			}

			fileList.Lock.Lock()
			defer fileList.Lock.Unlock()

			fileList.IdMap[f.Id] = &model.FileModel{
				Id:        f.Id,
				Name:      decodeName,
				CreatedAt: f.CreatedAt,
				Key:       f.Key,
				Bucket:    f.Bucket,
				Size:      f.Size,
				MimeType:  f.MimeType,
			}
		}(f)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		list = append(list, fileList.IdMap[id])
	}

	return list, count, nil
}
