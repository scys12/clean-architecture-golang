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
	CreateSession(echo.Context, *model.User) error
	Get(string) (Session, error)
	Set(string, Session) error
	Connect() redis.Conn
}

func (r *redisClient) Connect() redis.Conn {
	return r.conn
}

func (r *redisClient) CreateSession(ctx echo.Context, user *model.User) error {
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

func (r *redisClient) Set(sess_id string, session Session) error {
	s, err := json.Marshal(session)
	if err != nil {
		return err
	}
	if _, err = r.conn.Do("SET", sess_id, s); err != nil {
		return err
	}
	return nil
}
