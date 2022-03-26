package main

import "fmt"

var (
	sessionMrg SessionMgr
)

func Init(provider string, addr string, options ...string) (err error) {
	switch provider {
	case "memory":
		sessionMrg = NewMemorySessionMgr()
		break
	case "redis":
		sessionMrg = NewRedisSessionMgr()
		break
	default:
		fmt.Errorf("暂不支持")
	}
	sessionMrg.Init(addr, options...)
	return
}

func main() {
	Init("memory", "")
	session, _ := sessionMrg.CreateSession()
	session.Set("zhangsan", "adsa")
	fmt.Println(session.Get("zhangsan"))
}
