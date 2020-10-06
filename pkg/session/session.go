package session

import (
	"encoding/json"
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
	GetItems() (*[]model.Item, error)
	Set(FieldStore) error
	CheckLength(string, string) (int, error)
	Connect() redis.Conn
	Del(FieldStore) error
	GetSession(string) (*Session, error)
	InsertLatestItem(model.Item) error
	PopList(string, string) error
	TrimList(string, []string) error
	PushItem(string, model.Item) error
	UpdateItemList(string, string, int, model.Item) error
	UpdateItem(model.Item, model.Item) error
}

type FieldStore struct {
	TypeField string
	DataField []interface{}
}

const (
	sessionStr = "session"
	roleStr    = "role"
	itemsStr   = "items"
)

func (r *redisClient) Connect() redis.Conn {
	return r.conn
}

func (r *redisClient) UpdateItem(oldItem model.Item, updatedItem model.Item) error {
	idx, _ := r.CheckItemExist("LPOS", itemsStr, oldItem)
	if idx == -1 {
		return nil
	}
	err := r.UpdateItemList("LSET", itemsStr, idx, updatedItem)
	return err
}

func (r *redisClient) UpdateItemList(command, key string, index int, item model.Item) error {
	it, err := json.Marshal(item)
	if err != nil {
		return err
	}
	args := make([]interface{}, 3)
	args[0] = key
	args[1] = index
	args[2] = it
	err = r.Set(FieldStore{TypeField: command, DataField: args})
	return err
}

func (r *redisClient) CheckItemExist(command, key string, item model.Item) (int, error) {
	it, err := json.Marshal(item)
	if err != nil {
		return -1, err
	}
	args := make([]interface{}, 2)
	args[0] = key
	args[1] = it
	idx, err := redis.Int(r.Get(FieldStore{TypeField: command, DataField: args}))
	if err != nil {
		return -1, err
	}
	return idx, nil
}

func (r *redisClient) GetItems() (*[]model.Item, error) {
	limit := []string{itemsStr, "0", "-1"}
	args := make([]interface{}, len(limit))
	for idx, str := range limit {
		args[idx] = str
	}
	itemsBytes, err := redis.ByteSlices(r.Get(FieldStore{TypeField: "LRANGE", DataField: args}))
	if err != nil {
		return nil, err
	}
	items := []model.Item{}
	for _, item := range itemsBytes {
		var it model.Item
		if err = json.Unmarshal(item, &it); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return &items, nil
}

func (r *redisClient) InsertLatestItem(item model.Item) error {
	length, err := r.CheckLength("LLEN", itemsStr)
	if err != nil {
		return err
	}
	if length >= 10 {
		limit := []string{itemsStr, "0", "9"}
		err = r.TrimList("LTRIM", limit)
		if err != nil {
			return err
		}
		err = r.PopList("RPOP", itemsStr)
		if err != nil {
			return err
		}
	}
	err = r.PushItem("LPUSH", item)
	return err
}

func (r *redisClient) CheckLength(command, key string) (int, error) {
	length, err := redis.Int(r.Get(FieldStore{TypeField: command, DataField: redis.Args{}.Add(key)}))
	return length, err
}

func (r *redisClient) PopList(command, key string) error {
	args := make([]interface{}, 1)
	args[0] = key
	err := r.Set(FieldStore{TypeField: command, DataField: args})
	return err
}

func (r *redisClient) TrimList(command string, limit []string) error {
	args := make([]interface{}, len(limit))
	for idx, str := range limit {
		args[idx] = str
	}
	err := r.Set(FieldStore{TypeField: command, DataField: args})
	return err
}

func (r *redisClient) PushItem(command string, item model.Item) error {
	it, err := json.Marshal(item)
	if err != nil {
		return err
	}
	args := make([]interface{}, 2)
	args[0] = itemsStr
	args[1] = it
	if err = r.Set(FieldStore{TypeField: command, DataField: args}); err != nil {
		return err
	}
	return nil
}

func (r *redisClient) CreateSession(user *model.UserAuth) (string, error) {
	sessionID := uuid.New().String()
	sess := map[string]interface{}{
		sessionStr: sessionID,
		roleStr:    user.Role.Name,
	}
	config := map[string]string{"key": fmt.Sprintf("user:%v", user.ID.Hex()), "command": "HMSET"}
	if err := r.initializeAndSetDataHash(config, sess); err != nil {
		return "", err
	}
	auth := map[string]interface{}{
		sessionID: user.ID.Hex(),
	}
	config = map[string]string{"key": sessionStr, "command": "HSET"}
	if err := r.initializeAndSetDataHash(config, auth); err != nil {
		return "", err
	}
	return sessionID, nil
}

func (r *redisClient) initializeAndSetDataHash(config map[string]string, data map[string]interface{}) error {
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
	sess := []string{"session", sessionID}
	data := make([]interface{}, len(sess))
	for idx, str := range sess {
		data[idx] = str
	}
	userID, err := redis.String(r.Get(FieldStore{TypeField: "HGET", DataField: data}))
	if err != nil {
		return nil, err
	}
	user := fmt.Sprintf("user:%v", userID)
	data = make([]interface{}, 1)
	data[0] = user
	auth, err := redis.StringMap(r.Get(FieldStore{TypeField: "HGETALL", DataField: data}))
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
