package main

import (
	"fmt"
	"multimap"
	"slicemultimap"
)

func main() {
	applyMsgCh := make(chan *multimap.ClientsOpRecord, 1)
	notLeaderEventCh := make(chan interface{}, 1)
	noticeChs := &multimap.NoticeCh{applyMsgCh, notLeaderEventCh}

	m := slicemultimap.New()
	m.Put(3, noticeChs)
	fmt.Println(m.Get(3))
}
