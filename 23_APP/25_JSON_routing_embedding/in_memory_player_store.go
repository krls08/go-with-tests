package main

import "sync"

type InMemoryPlayerStore struct {
	scores map[string]int
	league []Player
	lock   sync.RWMutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		[]Player{},
		sync.RWMutex{},
	}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.scores[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.scores[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	return i.league
}
