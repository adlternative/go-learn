package main

import (
	"runtime"
	"strings"
)

type T struct {
	str string
}

func main() {
	// Allocate something on the heap.
	str := strings.Repeat("x", 1024)

	// Allocate a stack object that points to str.
	t := &T{str: str}

	// Pass a pointer to the stack object to g.
	s := g(t)

	// At this point, neither str or t are statically live.
	// They should only be held onto by the reference in s.

	if s.t.str[0] != 'x' {
		panic("bad")
	}
}

type S struct {
	t *T
}

//go:noinline
func g(t *T) (s S) {
	s = S{t: t}

	// s escapes, so it is moved to the heap.
	escape(&s)

	// t is stored in the heap copy of s, but it is a stack pointer.
	// We need that t pointer to keep str alive, and we have to find it during
	// stack scanning. Having it live in the heap means the stack scan of this
	// goroutine will not find the stack object at *t, and hence not keep
	// str alive.
	runtime.GC() // This collects str!

	// At this return the heap copy of s is copied back to the stack, so main can
	// read it (and find t, even though we never scanned *t).
	return
}

//go:noinline
func escape(x interface{}) {
	sink = x
	sink = nil
}

var sink interface{}
