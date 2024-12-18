package set

import "testing"

func TestSet(t *testing.T) {
    s := New()
    s.Add("apple")
    s.Add("banana")

    if !s.Has("apple") {
        t.Errorf("Expected 'apple' to exist in the set")
    }

    s.Remove("apple")
    if s.Has("apple") {
        t.Errorf("Expected 'apple' to be removed from the set")
    }

    elements := s.Elements()
    if len(elements) != 1 {
        t.Errorf("Expected 1 element, got %d", len(elements))
    }
}
