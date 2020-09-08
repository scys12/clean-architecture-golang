package session

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/scys12/clean-architecture-golang/model"

	"github.com/labstack/echo/v4"

	"github.com/google/uuid"

	"github.com/gomodule/redigo/redis"
)

type Session struct {
	UserID   primitive.ObjectID `json:"user_id"`
	UserRole string             `json:"user_role"`
}

type SessionStore interface {
	CreateSession() error
	Get(string) (Session, error)
	Set(string, Session) error
}

func (r *RedisClient) CreateSession(ctx echo.Context, user *model.User) error {
	sessionID := uuid.New().String()
	ctx.SetCookie(&http.Cookie{
		Name:     "sessionID",
		Value:    sessionID,
		HttpOnly: true,
		MaxAge:   86400 * 7,
	})
	ctx.Set("sessionID", sessionID)
	r.Set(sessionID, Session{UserID: user.ID, UserRole: user.Role.Name})
	return nil
}

func (r *RedisClient) Get(sess_id string) (Session, error) {
	var session Session
	id, err := redis.Bytes(r.Conn.Do("GET", sess_id))
	if err != nil {
		return session, err
	}
	if err = json.Unmarshal(id, &session); err != nil {
		return session, err
	}
	return session, nil
}

func (r *RedisClient) Set(sess_id string, session Session) error {
	s, err := json.Marshal(session)
	if err != nil {
		return err
	}
	if _, err = r.Conn.Do("SET", sess_id, s); err != nil {
		return err
	}
	return nil
}
