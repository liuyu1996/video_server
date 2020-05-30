package dbops

import (
	"video_server/scheduler/defs"
)

func AddVideoDeletionRecord(vid string) error {
	VideoDleRec := &defs.VideoDelRec{VideoId:vid}
	if err := dbConn.Create(VideoDleRec).Error; err != nil{
		return err
	}
	return nil
}
