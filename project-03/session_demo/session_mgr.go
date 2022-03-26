package main

type SessionMgr interface {
	// Init
	// @description 初始化
	// @param addr
	// @param options
	// @return error
	Init(addr string, options ...string) (err error)

	// CreateSession
	// @description 创建session
	// @return Session
	// @return error
	CreateSession() (session Session, err error)

	// Get
	// @description 通过sessionId获得session
	// @param sessionId
	// @return Session
	// @return error
	Get(sessionId string) (session Session, err error)
}
