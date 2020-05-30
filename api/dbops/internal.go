package dbops

import (
	"log"
	"sync"
	"video_server/api/defs"
)

func InsertSession(sid string, ttl string, userName string) error {
	session := &defs.Sessions{SessionId:sid, TTL:ttl, UserName:userName}
	if err := dbConn.Create(session).Error; err != nil{
		return err
	}
	return nil
}

func RetrieveSession(sid string) (*defs.Sessions, error) {
	session := &defs.Sessions{}
	if err := dbConn.Where("session_id=?", sid).Find(session).Error; err != nil{
		return nil, err
	}
	return session, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	rows, err := dbConn.Exec("SELECT * FROM sessions").Rows()
	if err != nil {
		return m, err
	}
	for rows.Next() {
		var id string
		var ttlstr string
		var userName string
		if er := rows.Scan(&id, &ttlstr, &userName); er != nil {
			log.Printf("retrive sessions error: %s", er)
			break
		}

		ss := &defs.Sessions{UserName: userName, TTL: ttlstr}
		m.Store(id, ss)
		log.Printf(" session id: %s, ttl: %d", id, ss.TTL)
	}
	return m, nil
}

func DeleteSessions(sid string) error {
	var session defs.Sessions
	if err := dbConn.Where("session_id=?", sid).Delete(&session).Error; err != nil {
		return err
	}
	return nil
}


