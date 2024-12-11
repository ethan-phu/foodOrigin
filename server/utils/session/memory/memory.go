// session的内存存储实现

package memory

import (
	"container/list"
	"errors"
	"knowFood/utils/session"
	"time"
)

// MemoryStore 每个session的内容的存储位置
type MemoryStore struct {
	// 会话id
	sid            string                      //session id
	lastAccessTime time.Time                   //session的最后访问时间
	value          map[interface{}]interface{} // 存储session内容的键值对容器
}

func (m *MemoryStore) Set(key string, value interface{}) error {
	if len(key) <= 0 {
		return errors.New("key cannot be empty!")
	}
	// 往map里面存
	m.value[key] = value
	return nil
}

func (m *MemoryStore) Get(key string) interface{} {
	if len(key) <= 0 {
		return nil
	}
	return m.value[key]
}

func (m *MemoryStore) Delete(key string) error {
	if len(key) <= 0 {
		return errors.New("key cannot empty when delete")
	}
	delete(m.value, key)
	return nil
}

func (m *MemoryStore) GetSessionId() string {
	return m.sid
}

// MemoryProvider 保存每一个session，实现的方式是，map+gclist（gc的列表）
type MemoryProvider struct {
	// session会话的元素节点
	sessions map[interface{}]*list.Element
	// 存放会话节点的列表
	gclist *list.List
}

func NewProvider() *MemoryProvider {
	return &MemoryProvider{
		sessions: make(map[interface{}]*list.Element, 0),
		gclist:   list.New(),
	}
}

func (p *MemoryProvider) SessionInit(sid string) (session.Session, error) {
	v := make(map[interface{}]interface{}, 0)
	s := &MemoryStore{
		sid:            sid,
		lastAccessTime: time.Now(),
		value:          v,
	}
	// 插入到链表尾巴
	element := p.gclist.PushBack(s)
	p.sessions[sid] = element
	return s, nil
}

func (p *MemoryProvider) SessionRead(sid string) (session.Session, error) {
	if v, ok := p.sessions[sid]; ok {
		sess := v.Value.(*MemoryStore)
		// 读到了旧的，更新会话的生命周期
		p.UpdateSessionLifeTime(v)
		return sess, nil
	} else {
		// 没读到就创建
		sess, err := p.SessionInit(sid)
		return sess, err
	}
}

func (p *MemoryProvider) SessionDestroy(sid string) error {
	if v, ok := p.sessions[sid]; ok {
		delete(p.sessions, sid)
		p.gclist.Remove(v)
	}
	return nil
}

// session的gc处理
func (p *MemoryProvider) SessionGC(maxLifeTime int64) {
	for {
		// 获得链表的最后一个元素
		element := p.gclist.Back()
		if element == nil {
			break
		}
		// session的最后访问时间+session的最大生命周期如果小于当前的时间，那么这个会话就过期了，就需要把内存里面的session删除掉
		if element.Value.(*MemoryStore).lastAccessTime.Unix()+maxLifeTime < time.Now().Unix() {
			p.gclist.Remove(element)
			delete(p.sessions, element.Value.(*MemoryStore).sid)
		} else {
			break
		}
	}
}

// UpdateSessionLifeTime 每次读session更新session最后访问时间，更新会话时间，只要一直来就会保持会话的有效性
func (p *MemoryProvider) UpdateSessionLifeTime(e *list.Element) {
	e.Value.(*MemoryStore).lastAccessTime = time.Now()
	// 更新了会话，将session节点移动到链表的头部
	p.gclist.MoveToFront(e)
}
