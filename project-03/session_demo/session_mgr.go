package session_demo

type SessionMgr interface {
	// Init
	// @description 初始化
	// @param addr
	// @param options
	// @return error
	Init(addr string, options ...string) error

	// CreateSession
	// @description 创建session
	// @return Session
	// @return error
	CreateSession() (Session, error)

	// Get
	// @description 通过sessionId获得session
	// @param sessionId
	// @return Session
	// @return error
	Get(sessionId string) (Session, error)
}
