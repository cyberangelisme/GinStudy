package session

import (
	"errors"
	"sync"
)

//MemorySession
//构造函数获取对象
//结构体的方法组
type MemorySession struct {
	sessionId string
	data      map[string]interface{} //存k-v键值对
	rwlock    sync.RWMutex           //读写锁
	maxAge    int64
}

//构造函数
func NewMemorySession() *MemorySession {
	s := &MemorySession{

		data:   make(map[string]interface{}, 16),
		maxAge: 30,
	}
	return s
}

//输入键，获取对应session的————值和error
func (m *MemorySession) Set(key string, value interface{}) (err error) {
	//加读写锁，在该routine读锁占用时，禁止其他routine的写操作，但不禁止读操作
	//写锁会阻止其他gorotine不论读或者写进来
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	//设置值
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	//没取到是bool，false
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

func (m *MemorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return
}

func (m *MemorySession) Save() (err error) {
	return
}
