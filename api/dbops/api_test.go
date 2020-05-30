package dbops

import (
	"strconv"
	"testing"
	"time"
)

var videoId string

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}

func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestUserWorkFlow(t *testing.T)  {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func TestVideoWorkFlow(t *testing.T)  {
	t.Run("prepareUser", testAddUser)
	t.Run("AddVideo", TestAddVideo)
	t.Run("GetVideo", TestGetVideo)
	t.Run("DeleteVideo", TestDeleteVideo)
	t.Run("RegetVideo", testRegetVideo)
}

func testAddUser(t *testing.T)  {
	err := AddUserCredential("liuyu", "123456")
	if err != nil {
		t.Errorf("add user failed: %v", err)
	}
}

func testGetUser(t *testing.T)  {
	pwd, err := GetUserCredential("liuyu")
	if err != nil || pwd != "123456" {
		t.Errorf("get user failed: %v", err)
	}
}

func testDeleteUser(t *testing.T)  {
	err := DeleteUserCredential("liuyu", "1234566")
	if err != nil {
		t.Errorf("delete user failed:%v", err)
	}
}

func testRegetUser(t *testing.T)  {
	pwd, _ := GetUserCredential("liuyu")

	if pwd != "" {
		t.Errorf("delete user failed")
	}
}

func TestAddVideo(t *testing.T)  {
	video, err := AddVideo(1, "1111")
	if err != nil {
		t.Errorf("add video failed: %v", err)
	}
	videoId = video.Id
}

func TestGetVideo(t *testing.T)  {
	video, err := GetVideo(videoId)
	if err != nil || video == nil {
		t.Errorf("get video failed: %v", err)
	}
}

func TestDeleteVideo(t *testing.T)  {
	err := DeleteVideo(videoId)
	if err != nil {
		t.Errorf("delete video failed:%v", err)
	}
}

func testRegetVideo(t *testing.T)  {
	video, _ := GetVideo(videoId)
	if video != nil {
		t.Error("delelte video failed")
	}
}

func TestComments(t *testing.T)  {
	clearTables()
	t.Run("addc", testAddComments)
	t.Run("listc", testListComments)
}

func testAddComments(t *testing.T)  {
	vid := "111"
	aid := 1
	content := "123456"

	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("add comment failed: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "111"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	_, err := ListComments(vid,from, to)
	if err != nil {
		t.Errorf("list comment failed: %v", err)
	}
}
