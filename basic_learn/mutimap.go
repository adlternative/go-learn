// Package multimap provides an abstract MultiMap interface.
//
// Multimap is a collection that maps keys to values, similar to map.
// However, each key may be associated with multiple values.
//
// You can visualize the contents of a multimap either as a map from keys to nonempty collections of values:
//    - a --> 1, 2
//    - b --> 3
// ... or a single "flattened" collection of key-value pairs.
//    - a --> 1
//    - a --> 2
//    - b --> 3
//
// Similar to a map, operations associated with this data type allow:
// - the addition of a pair to the collection
// - the removal of a pair from the collection
// - the lookup of a value associated with a particular key
// - the lookup whether a key, value or key/value pair exists in this data type.
package main

type Op struct {
	Key   string
	Value string
	Op    string // "Put" or "Append" or "Get"

	ClientId int64
	Seq      int32
}

/* 需要持久化 */
type ClientsOpRecord struct {
	/* 操作 */
	Op Op
	/* 结果 */
	ResultValue string
	Error       string
}

type NoticeCh struct {
	ApplyMsgChs       chan *ClientsOpRecord
	NotLeaderEventChs chan interface{}
}

type ActiveClient struct {
	ClientId    int64
	NoticeChMap *slicemultimap.MultiMap
}

// Entry represents a key/value pair inside a multimap.
type Entry struct {
	Key   int32
	Value *NoticeCh
}

// MultiMap interface that all multimaps implement.
type MultiMap interface {
	Get(key interface{}) (value []interface{}, found bool)

	Put(key interface{}, value interface{})
	PutAll(key interface{}, value []interface{})

	Remove(key interface{}, value interface{})
	RemoveAll(key interface{})

	Contains(key interface{}, value interface{}) bool
	ContainsKey(key interface{}) bool
	ContainsValue(value interface{}) bool

	Entries() []Entry
	Keys() []interface{}
	KeySet() []interface{}
	Values() []interface{}

	Clear()
	Empty() bool
	Size() int
}
