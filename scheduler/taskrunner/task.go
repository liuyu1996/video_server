package taskrunner

import (
	"errors"
	"log"
	"os"
	"sync"
	"video_server/scheduler/dbops"
)

func deleteVideoFile(vid string) error {
	err := os.Remove(VIDEO_DIR + vid)
	if err !=nil && !os.IsNotExist(err) {
		log.Printf("Delete Video Failed")
		return err
	}
	return nil

	//ossfn := "videos/" + vid
	//bn := ""
	//ok := aliyunoss.DeleteObject(vid, bn)
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Println(err)
		return err
	}
	if len(res) == 0 {
		return errors.New("All task finished")
	}

	for _, id := range res{
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
	forloop:
		for  {
			select {
			case vid := <- dc:
				go func(id interface{}) {
					if err := deleteVideoFile(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
					if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
				}(vid)
			default:
				break forloop
			}
		}
		errMap.Range(func(key, value interface{}) bool {
			err = value.(error)
			if err != nil {
				return false
			}
			return true
		})
	return err
}
