package sequence

import (
	"sync"
)

// just for testing
type MemorySequencer struct {
	counter      uint64
	sequenceLock sync.Mutex
}

func NewMemorySequencer() (m *MemorySequencer) {
	m = &MemorySequencer{counter: 1}
	return
}

func (m *MemorySequencer) NextFileId(count int) (uint64, int) {
	m.sequenceLock.Lock()
	defer m.sequenceLock.Unlock()
	ret := m.counter
	m.counter += uint64(count)
	return ret, count
}

func (m *MemorySequencer) SetMax(seenValue uint64) {
	m.sequenceLock.Lock()
	defer m.sequenceLock.Unlock()
	if m.counter <= seenValue {
		m.counter = seenValue + 1
	}
}

func (m *MemorySequencer) Peek() uint64 {
	return m.counter
}
