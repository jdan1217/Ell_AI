package traffic

import (
	"testing"
)

func TestHistoryAdd(t *testing.T) {
	entry := TrafficEntry{
		ID: "test-123",
	}
	history := &History{}
	history.Add(entry)

	entries := history.All()

	if len(entries) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(entries))
	}

	if entries[0].ID != entry.ID {
		t.Errorf("Expected ID %s, got %s", entry.ID, entries[0].ID)
	}
}

func TestHistoryAll(t *testing.T) {
	history := &History{}

	history.Add(TrafficEntry{ID: "first"})
	history.Add(TrafficEntry{ID: "second"})
	history.Add(TrafficEntry{ID: "third"})

	entries := history.All()

	if len(entries) != 3 {
		t.Fatalf("Expected 3 entries, got %d", len(entries))
	}

	expected := []string{"first", "second", "third"}

	for i, entry := range entries {
		if entry.ID != expected[i] {
			t.Errorf("Expected entry %s at index %d, got %s", expected[i], i, entry.ID)
		}
	}
}

func TestHistoryGetByID(t *testing.T) {
	history := &History{}

	history.Add(TrafficEntry{ID: "first"})
	history.Add(TrafficEntry{ID: "second"})
	history.Add(TrafficEntry{ID: "third"})

	entry, found := history.GetByID("second")

	if !found {
		t.Fatal("Expected to find entry with ID second")
	}

	if entry.ID != "second" {
		t.Errorf("Expected ID second, got %s", entry.ID)
	}
}
