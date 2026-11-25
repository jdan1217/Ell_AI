package traffic

import (
	"sync"
)

type History struct {
	mu      sync.RWMutex
	entries []TrafficEntry
}

func (h *History) Add(entry TrafficEntry) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.entries = append(h.entries, entry)
}

func (h *History) All() []TrafficEntry {
	h.mu.RLock()
	defer h.mu.RUnlock()
	entries := make([]TrafficEntry, len(h.entries))
	copy(entries, h.entries)
	return entries
}

func (h *History) GetByID(id string) (TrafficEntry, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, entry := range h.entries {
		if entry.ID == id {
			return entry, true
		}
	}
	return TrafficEntry{}, false
}

func (h *History) Paginated(limit, offset int) []TrafficEntry {
	h.mu.RLock()
	defer h.mu.RLock()
	if offset >= len(h.entries) {
		return []TrafficEntry{}
	}
	end := offset + limit
	if end > len(h.entries) {
		end = len(h.entries)
	}
	entries := make([]TrafficEntry, end-offset)
	copy(entries, h.entries[offset:end])
	return entries
}
