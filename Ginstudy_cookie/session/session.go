package session

//定义Session接口，管理单个session中的内容
type Session interface {
	Set(key string, value interface{}) error //设置session
	Get(key string) (interface{}, error)     //获取session
	Del(key string) error                    //删除session
	Save() error                             //保存
}

// type Session interface {
// 	Set(key, value interface{}) error // set session value
// 	Get(key interface{}) interface{}  // get session value
// 	Delete(key interface{}) error     // delete session value
// 	SessionID() string                // back current sessionID
// }
