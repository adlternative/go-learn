package memo4

import "sync"

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	f     Func
	mu    sync.Mutex        // guards cache
	cache map[string]*entry //
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// NOTE: not concurrency-safe!
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	/* 注意这个缓存后map中保存了key对应的entry e,里面的通道ready是关闭状态,
	读取它的那些家伙能不阻塞返回 */
	e := memo.cache[key]
	/* 如果没有拿到缓存 */
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready /* 可能阻塞读,如果管道关闭了就是非阻塞了 */
	}

	return e.res.value, e.res.err
}
