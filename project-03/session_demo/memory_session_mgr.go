package session_demo

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

// MemorySessionMgr
// @description 数据结构如下 sessionId,Map<id,value>
type MemorySessionMgr struct {
	sessionMap map[string]*Session
	rwLock     sync.RWMutex
}

// NewMemorySessionMgr
// @description 构造函数返回*MemorySessionMgr
// @return *MemorySessionMgr
func NewMemorySessionMgr() *MemorySessionMgr {
	m := &MemorySessionMgr{
		sessionMap: make(map[string]*Session, 1024),
	}
	return m
}

func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {

	return
}

func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	id := uuid.NewV4()

	sessionId := id.String()
	session = NewMemorySession(sessionId)

	return
}
