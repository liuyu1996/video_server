package defs




type Users struct {
	Id int
	Username string `json:"user_name" gorm:"column:user_name"`
	Pwd string `json:"pwd" gorm:"column:pwd"`
}

type VideoInfo struct {
	Id string `json:"id"`
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comments struct {
	Id       string `json:"id"`
	VideoId  string `json:"video_id"`
	AuthorId int `json:"author_id"`
	Content  string `json:"content"`
}

type CommentData struct {
	Id       string `json:"id"`
	VideoId  string `json:"video_id"`
	UserName string `json:"user_name"`
	Content  string `json:"content"`
}

type Sessions struct {
	SessionId string `json:"session_id"`
	UserName string `json:"user_name"`
	TTL string `json:"ttl"`
}

//response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}