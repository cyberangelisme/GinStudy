package session

//管理所有session的接口
type SessionMgr interface {
	SessionInit(addr string, options ...string) error
	CreateSession() (Session, error)
	SessionGet(sessionId string) (Session, error)
	SessionDestory(sessionId string) error
	SessionGC(maxLifeTime int64)
}

// type Provider interface {
// 	SessionInit(sid string) (Session, error)
// 	SessionRead(sid string) (Session, error)
// 	SessionDestroy(sid string) error
// 	SessionGC(maxLifeTime int64)
// }
