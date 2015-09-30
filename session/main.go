package main

import (
	"fmt"
	"http"
	"io"
)

type Manager struct {
	cookieName  string     //cookieName
	lock        sync.Mutex //锁
	provider    provider   //procvider接口，表征session管理器底层存储结构
	maxlifetime int64      //最大有效时间
}

/*全局Session管理器*/

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) { //返回manager的指针
	provider, ok := providers[provideName] //根据providerName在providers中拿到对应的provider
	if !ok {
		return nil, fmt.Errorf("session: unknow provide %q (forgotten import?)", provideName) //找不到provider,报错提示
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil //若找到，则返回一个Manager指针
}

var gloablSession *session.Manager

func init() { //实例化session管理器
	gloablSession = NewManager("memory", "gosessionid", 3600)
}

type Provider interface { //定义Provider接口，负责管理底层存储。抽象底层存储
	SessionInit(sid string) (Session, error) //初始化session，返回session变量
	SessionRead(sid string) (Session, error) //返回sid代表的session变量，若不存在则取创建一个
	SessionDestory(sid string) error         //销毁sid对应的session变量
	SessionGC(maxLifeTime int64)             //根据maxlifetime删除过期session
}

type Session interface { //定义session接口，抽象不同的Provider
	Set(key, value interface{}) error //设置session值
	Get(key interface{}) interface{}  //获取session值
	Delete(key interface{}) error     //删除session值
	SessionID() string                //获取当前sessionID
}

var providers = make(map[string]Provider) //定义一个providers的map,供全局session使用

// 注册函数，通过注册使得不同的session provide可以通过providerName被使用

func Register(name string, provider Provider) {
	if driver == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := providers[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	providers[name] = provider
}

// 全局唯一的session ID

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEcoding.EncodeToString(b)
}

// 为每个用户关联一个session

//模拟login
func login(w http.responseWriter, r *http.Request) {
	sess := gloablSession.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}
