package utils

import "sync"

type SyncChannel struct {
	C    chan bool
	Once sync.Once
}

func NewSyncChannel() *SyncChannel {
	return &SyncChannel{C: make(chan bool)}
}
