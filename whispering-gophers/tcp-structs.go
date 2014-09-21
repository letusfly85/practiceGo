package myserver

import "sync"

type Peers struct {
	m  map[string]chan<- Message
	mu sync.RWMutex
}

type Message string

type Site struct {
	Addr    string
	Message string
	URL     string
}
