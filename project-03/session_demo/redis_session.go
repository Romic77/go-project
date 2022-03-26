package main

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"sync"
)

type RedisSession struct {
	sessionId  string
	pool       *redis.Pool
	sessionMap map[string]interface{}
	//读写锁
	rwLock sync.RWMutex
	//记录内存中map是否被操作
	flag int
}

const (
	//内存数据没变化
	SessionFlagNone = iota
	//内存数据有修改
	SessionFlagModify
)

// NewRedisSession
// @description 构造函数返返回RedisSession指针
// @param sessionId string
// @param pool *redis.Pool
// @return *RedisSession
func NewRedisSession(sessionId string, pool *redis.Pool) Session {
	return &RedisSession{
		sessionId:  sessionId,
		pool:       pool,
		sessionMap: make(map[string]interface{}, 16),
		flag:       SessionFlagNone,
	}
}

// Set
// @description 设置值
// @receiver r RedisSession
// @param key string
// @param value interface{}
// @return err error
func (r *RedisSession) Set(key string, value interface{}) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	//设置值
	r.sessionMap[key] = value
	r.flag = SessionFlagModify
	return
}

func (r *RedisSession) Save() (err error) {
	//加锁
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	//如果数据没有变化，不需要存
	if r.flag != SessionFlagModify {
		return
	}

	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return err
	}

	//获取redis链接
	r.pool.Get().Do("SET", r.sessionId, string(data))
	return
}

func (r *RedisSession) Get(key string) (value interface{}, err error) {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	value, ok := r.sessionMap[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

func (r *RedisSession) GetFromRedis(key string) (value interface{}, err error) {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	reply, err := r.pool.Get().Do("GET", key)
	if err != nil {
		return
	}
	data, _ := redis.String(reply, err)
	json.Unmarshal([]byte(data), &r.sessionMap)
	return
}

func (r *RedisSession) Delete(key string) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	delete(r.sessionMap, key)
	r.flag = SessionFlagModify
	return
}
