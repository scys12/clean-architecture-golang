package session

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scys12/clean-architecture-golang/usecase/user"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/google/uuid"

	"github.com/gomodule/redigo/redis"
)

type Session struct {
	UserID   primitive.ObjectID `json:"user_id" redis:"user_id"`
	UserRole string             `json:"user_role" redis:"user_role"`
}

type SessionStore interface {
	CreateSession(*user.Response) (*http.Cookie, error)
	Get(string) (Session, error)
	Set(fieldStore) error
	Connect() redis.Conn
	Del(key string) error
}

type fieldStore struct {
	typeField string
	dataField []interface{}
}

const (
	session_str   = "session"
	role_str      = "role"
	sessionID_str = "sessionID"
)

func (r *redisClient) Connect() redis.Conn {
	return r.conn
}

func (r *redisClient) CreateSession(user *user.Response) (*http.Cookie, error) {
	sessionID := uuid.New().String()
	sess := map[string]interface{}{
		session_str: sessionID,
		role_str:    user.RoleName,
	}
	data := redis.Args{}.Add(fmt.Sprintf("user:%v", user.ID)).AddFlat(sess)
	r.Set(fieldStore{typeField: "HMSET", dataField: 'data'})
	sess = map[string]interface{}{
		sessionID: user.ID,
	}
	data = redis.Args{}.Add(session_str).AddFlat(sess)
	r.Set(fieldStore{typeField: "HSET", dataField: data})

	cookie := &http.Cookie{
		Name:     sessionID_str,
		Value:    sessionID,
		HttpOnly: true,
		MaxAge:   86400 * 7,
	}

	return cookie, nil
}

func (r *redisClient) Get(sess_id string) (Session, error) {
	var session Session
	redis.Strings
	id, err := redis.Bytes(r.conn.Do("GET", sess_id))
	if err != nil {
		return session, err
	}
	if err = json.Unmarshal(id, &session); err != nil {
		return session, err
	}
	return session, nil
}

func (r *redisClient) Set(dataStore fieldStore) error {
	if _, err := r.conn.Do(dataStore.typeField, dataStore.dataField...); err != nil {
		return err
	}
	return nil
}

func (r *redisClient) Del(key string) error {
	if _, err := r.conn.Do("DEL", key); err != nil {
		return err
	}
	return nil
}
