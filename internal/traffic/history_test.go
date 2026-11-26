package traffic

import (
	"fmt"
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

func TestHistoryGetByIDNotFound(t *testing.T) {
	history := &History{}

	history.Add(TrafficEntry{ID: "first"})

	entry, found := history.GetByID("missing")

	if found {
		t.Fatal("Expected not to find entry")
	}

	if entry.ID != "" {
		t.Errorf("Expected empty entry, got ID %s", entry.ID)
	}
}

func TestFirstPageEntries(t *testing.T) {
	history := &History{}

	for i := 1; i <= 5; i++ {
		history.Add(TrafficEntry{ID: fmt.Sprintf("%d", i)})
	}

	entries := history.Paginated(2, 0)

	if len(entries) != 2 {
		t.Fatalf("Expected 2 entries, got %d", len(entries))
	}

	if entries[0].ID != "1" || entries[1].ID != "2" {
		t.Fatalf("Unexpected entries: %+v", entries)
	}
}

func TestTwoPagesEntries(t *testing.T) {
	history := &History{}

	for i := 1; i <= 5; i++ {
		history.Add(TrafficEntry{ID: fmt.Sprintf("%d", i)})
	}

	firstPage := history.Paginated(2, 0)
	secondPage := history.Paginated(2, 2)

	if len(firstPage) != 2 || len(secondPage) != 2 {
		t.Fatalf("Expected 2 entries per page, got %d and %d",
			len(firstPage), len(secondPage))
	}

	if firstPage[0].ID != "1" || firstPage[1].ID != "2" {
		t.Fatalf("Unexpected first page: %+v", firstPage)
	}

	if secondPage[0].ID != "3" || secondPage[1].ID != "4" {
		t.Fatalf("Unexpected second page: %+v", secondPage)
	}
}

func TestOffsetTooFar(t *testing.T) {
	history := &History{}

	for i := 1; i <= 5; i++ {
		history.Add(TrafficEntry{ID: fmt.Sprintf("%d", i)})
	}

	entries := history.Paginated(2, 10)

	if len(entries) != 0 {
		t.Fatalf("Expected 0 entries, got %d", len(entries))
	}
}

func TestLimitLargerThanEntries(t *testing.T) {
	history := &History{}

	for i := 1; i <= 3; i++ {
		history.Add(TrafficEntry{ID: fmt.Sprintf("%d", i)})
	}

	entries := history.Paginated(10, 1)

	if len(entries) != 2 {
		t.Fatalf("Expected 2 entries, got %d", len(entries))
	}

	if entries[0].ID != "2" || entries[1].ID != "3" {
		t.Fatalf("Unexpected entries: %+v", entries)
	}
}
