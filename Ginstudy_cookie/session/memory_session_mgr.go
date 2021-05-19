package session

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MemorySessionMgr struct {
	sessionMap  map[string]Session //保存每一个session
	rwlock      sync.RWMutex
	sessionMgr  SessionMgr
	maxLifeTime int64
}

//构造函数初始化管理器
func NewMemorySessionMgr() *MemorySessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024), //最大session数
	}
	return sr
}

//Init一个session,返回session接口变量
func (sm *MemorySessionMgr) SessionInit(sid string) (Session, error) {
	sm.rwlock.Lock()
	defer sm.rwlock.Unlock()
	newSession := NewMemorySession()
	newSession.sessionId = sid

	//将session存在管理器中（内存map）
	sm.sessionMap[sid] = newSession
	//返回接口session的变量时,要保证MemorySession这个结构体实现了session中所有接口
	return newSession, nil
	//随后可以利用返回的接口变量,调用接口中的方法-----newsess.Get()
}

//CURD
//创建一个新session，返回该session和err
func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	//go get github.com/satori/go.uuid
	//uuid(根据机器和时空特性的唯一性id)作为sessionId
	id := uuid.NewV4()

	//id转为string类型
	sessionId := id.String()
	//创建一个新session
	session = NewMemorySession(sessionId)
	return
}
func (s *MemorySessionMgr) Get() (session Session, err error) {

}
