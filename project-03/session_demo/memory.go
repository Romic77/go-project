package main

import (
	"errors"
	"sync"
)

// MemorySession
// @description 内存会话
type MemorySession struct {
	sessionId string
	// 数据
	data map[string]interface{}
	//读写锁
	rwLock sync.RWMutex
}

// NewMemorySession
// @description 创建MemorySession
// @param id string
// @return *MemorySession
func NewMemorySession(id string) Session {
	return &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
}

// Set
// @description 设置内存session的值
// @receiver m *MemorySession
// @param key string
// @param value interface{}
// @return err error
func (m *MemorySession) Set(key string, value interface{}) (err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	//设置值
	m.data[key] = value
	return
}

// Get
// @description 通过key获得session
// @receiver m *MemorySession
// @param key string
// @return interface{}
// @return error
func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

// Delete
// @description 根据key删除map中的元素
// @receiver m *MemorySession
// @param key string
// @return err error
func (m *MemorySession) Delete(key string) (err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	delete(m.data, key)
	return
}

// Save
// @description 保存 不知道是干嘛的
// @receiver m *MemorySession
// @return err error
func (m *MemorySession) Save() (err error) {
	return
}
