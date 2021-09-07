package utils

import "sync"

// SyncChannel is a structure for util once logic.
type SyncChannel struct {
	C    chan bool
	Once sync.Once
}

// NewSyncChannel returns SyncChannel pointer.
func NewSyncChannel() *SyncChannel {
	return &SyncChannel{C: make(chan bool)}
}
