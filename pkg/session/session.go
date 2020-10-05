package session

import (
	"fmt"

	"github.com/scys12/clean-architecture-golang/model"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/google/uuid"

	"github.com/gomodule/redigo/redis"
)

type Session struct {
	SessionID string             `json:"session_id" redis:"session_id"`
	UserID    primitive.ObjectID `json:"user_id" redis:"user_id"`
	UserRole  string             `json:"user_role" redis:"user_role"`
}

type SessionStore interface {
	CreateSession(*model.UserAuth) (string, error)
	Get(FieldStore) (interface{}, error)
	Set(FieldStore) error
	Connect() redis.Conn
	Del(FieldStore) error
	GetSession(string) (*Session, error)
}

type FieldStore struct {
	TypeField string
	DataField []interface{}
}

const (
	sessionStr = "session"
	roleStr    = "role"
)

func (r *redisClient) Connect() redis.Conn {
	return r.conn
}

func (r *redisClient) CreateSession(user *model.UserAuth) (string, error) {
	sessionID := uuid.New().String()
	sess := map[string]interface{}{
		sessionStr: sessionID,
		roleStr:    user.Role.Name,
	}
	config := map[string]string{"key": fmt.Sprintf("user:%v", user.ID.Hex()), "command": "HMSET"}
	if err := r.initializeAndSetData(config, sess); err != nil {
		return "", err
	}
	auth := map[string]interface{}{
		sessionID: user.ID.Hex(),
	}
	config = map[string]string{"key": sessionStr, "command": "HSET"}
	if err := r.initializeAndSetData(config, auth); err != nil {
		return "", err
	}
	return sessionID, nil
}

func (r *redisClient) initializeAndSetData(config map[string]string, data map[string]interface{}) error {
	args := redis.Args{}.Add(config["key"]).AddFlat(data)
	err := r.Set(FieldStore{TypeField: config["command"], DataField: args})
	return err
}

func (r *redisClient) Get(dataStore FieldStore) (interface{}, error) {
	data, err := r.conn.Do(dataStore.TypeField, dataStore.DataField...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *redisClient) GetSession(sessionID string) (*Session, error) {
	userID, err := redis.String(r.conn.Do("HGET", "session", sessionID))
	if err != nil {
		return nil, err
	}
	auth, err := redis.StringMap(r.conn.Do("HGETALL", fmt.Sprintf("user:%v", userID)))
	if err != nil {
		return nil, err
	}
	newUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	return &Session{
		UserID:    newUserID,
		SessionID: auth[sessionStr],
		UserRole:  auth[roleStr],
	}, nil
}

func (r *redisClient) Set(dataStore FieldStore) error {
	if _, err := r.conn.Do(dataStore.TypeField, dataStore.DataField...); err != nil {
		return err
	}
	return nil
}

func (r *redisClient) Del(dataStore FieldStore) error {
	if _, err := r.conn.Do(dataStore.TypeField, dataStore.DataField...); err != nil {
		return err
	}
	return nil
}
