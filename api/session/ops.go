package session

import (
	"strconv"
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
)


var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB()  {
	m ,err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	m.Range(func(key, value interface{}) bool {
		ss := value.(*defs.Sessions)
		sessionMap.Store(key,ss)
		return true
	})
}

func GenerateNewSessionId(userName string) string {
	id, _ := utils.NewUUID()
	createTime := time.Now().UnixNano()/1000000
	ttl := createTime + 30 * 60 * 1000
	ttlStr := strconv.FormatInt(ttl, 10)
	ss := &defs.Sessions{UserName:userName, TTL:ttlStr}
	sessionMap.Store(id, ss)
	_ = dbops.InsertSession(id, ttlStr, userName)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		t := time.Now().UnixNano()/1000000
		ttl, err := strconv.ParseInt(ss.(*defs.Sessions).TTL, 10, 64)
		if err != nil {
			return "", true
		}
		if ttl < t {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.Sessions).UserName, false
	}
	return "", true
}

func deleteExpiredSession(sid string)  {
	sessionMap.Delete(sid)
	_ = dbops.DeleteSessions(sid)
}