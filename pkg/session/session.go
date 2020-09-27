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
	UserID   primitive.ObjectID `json:"user_id"`
	UserRole string             `json:"user_role"`
}

type SessionStore interface {
	CreateSession(*user.Response) error
	Get(string) (Session, error)
	Set(string, string, Session) error
	Connect() redis.Conn
	Del(key string) error
}

func (r *redisClient) Connect() redis.Conn {
	return r.conn
}

func (r *redisClient) CreateSession(user *user.Response) error {
	sessionID := uuid.New().String()
	cookie := &http.Cookie{
		Name:     "sessionID",
		Value:    sessionID,
		HttpOnly: true,
		MaxAge:   86400 * 7,
	}
	r.Set("HSET", fmt.Sprintf("user:%v", user.ID), "session")
	r.Set("HSET", sessionID, user.ID)
	return nil
}

func (r *redisClient) Get(sess_id string) (Session, error) {
	var session Session
	id, err := redis.Bytes(r.conn.Do("GET", sess_id))
	if err != nil {
		return session, err
	}
	if err = json.Unmarshal(id, &session); err != nil {
		return session, err
	}
	return session, nil
}

func (r *redisClient) Set(type_data, field string, data ...interface{}) error {
	s, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err = r.conn.Do(type_data, sess_id, s); err != nil {
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
