package dbops

import (
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

func AddUserCredential(userName string, pwd string) error  {
	hashPwd, err := utils.PasswordHash(pwd)
	if err != nil {
		return err
	}
	user := &defs.Users{Username:userName, Pwd:hashPwd}
	if err := dbConn.Create(user).Error; err != nil{
		return err
	}
	return nil
}

func GetUser(userName string) (*defs.Users, error) {
	user := &defs.Users{}
	if err := dbConn.Where("user_name=?", userName).First(user).Error; err != nil{
		return nil, err
	}
	return user, nil
}

func GetUserCredential(userName string) (string ,error) {
	var user defs.Users
	if err := dbConn.Select("pwd").Where("user_name=?", userName).Find(&user).Error; err != nil{
		return "", err
	}
	return user.Pwd, nil
}

func DeleteUserCredential(userName string, pwd string) error {
	user := &defs.Users{Username:userName,Pwd:pwd}
	if err := dbConn.Delete(user).Error; err != nil{
		return err
	}
	return nil
}

func AddVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	uid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	timeNow := time.Now()
	ctime := timeNow.Format("Jan 02 2006, 15:04:05")
	video := &defs.VideoInfo{Id:uid, AuthorId:aid, Name:name, DisplayCtime:ctime}
	if err := dbConn.Create(video).Error; err != nil{
		return nil, err
	}
	return video, nil
}

func GetVideo(vid string) (*defs.VideoInfo, error) {
	video := &defs.VideoInfo{}
	if err := dbConn.Where("id=?",vid).Find(video).Error; err !=nil {
		return nil, err
	}
	return video, nil
}

func GetAllVideo(uname string, from, to int) ([] *defs.VideoInfo, error) {
	var videoList [] *defs.VideoInfo
	if err := dbConn.Table("video_info").Select("*").Joins("inner join users on users.id=video_info.author_id").Where("users.user_name=?" +
		"and video_info.create_time>FROM_UNIXTIME(?) and video_info.create_time<=FROM_UNIXTIME(?)",uname, from, to).Find(videoList).Error; err != nil{
		return nil, err
	}
	return videoList, nil
}

func DeleteVideo(vid string) error {
	video := &defs.VideoInfo{}
	if err := dbConn.Where("id=?", vid).Delete(video).Error; err != nil{
		return err
	}
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	comment := &defs.Comments{Id:id, VideoId:vid, AuthorId:aid, Content:content}
	if err := dbConn.Create(comment).Error; err != nil{
		return err
	}
	return nil
}

func ListComments(vid string, from, to int) ([] *defs.CommentData, error) {
	var commentdata  []*defs.CommentData
	rows, err := dbConn.Table("comments").Select("comments.id, users.user_name, comments.content").Joins("inner join users on comments.author_id=users.id").Where(
		"comments.video_id=? and comments.time>FROM_UNIXTIME(?) and comments.time<=FROM_UNIXTIME(?)",vid, from, to).Order("comments.time desc").Rows()
	if err != nil {
		return commentdata, err
	}
	for rows.Next() {
		var commentId, userName, content string
		if err := rows.Scan(&commentId, userName, content); err !=nil {
			return commentdata, err
		}
		c := &defs.CommentData{Id:commentId, VideoId:vid, UserName:userName, Content:content}
		commentdata = append(commentdata, c)
	}
	return commentdata, nil
}
