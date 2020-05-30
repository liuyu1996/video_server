package dbops

import (
	"video_server/api/defs"
)

func ReadVideoDeletionRecord(count int) ([]string, error) {
	var ids[]string
	if err := dbConn.Table("video_del_rec").Select("video_id").Find(&ids).Limit(count).Error; err !=nil{
		return ids, err
	}
	return ids, nil
}

func DelVideoDeletionRecord(vid string) error {
	if err := dbConn.Delete(defs.VideoInfo{}).Where("video_id=?", vid).Error; err != nil{
		return err
	}
	return nil
}



