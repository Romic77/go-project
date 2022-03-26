package main

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	addr       string
	password   string
	pool       *redis.Pool
	rwLock     sync.RWMutex
	sessionMap map[string]Session
}

// NewRedisSessionMgr
// @description 构造函数返回SessionMgr,实现这个接口的所有方法
// @param addr string
// @param password string
// @return *RedisSessionMgr
func NewRedisSessionMgr() SessionMgr {
	return &RedisSessionMgr{
		sessionMap: make(map[string]Session, 32),
	}
}

func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	if len(options) > 0 {
		r.password = options[0]
	}
	//创建连接池
	r.pool = myPool(addr, r.password)
	r.addr = addr
	return
}

func myPool(addr string, password string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			conn.Do("AUTH", password)
			return conn, err
		},
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 60 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return err
			}
			return err
		},
	}
}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	id := uuid.NewV4()

	sessionId := id.String()
	session = NewMemorySession(sessionId)
	r.sessionMap[sessionId] = session
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
