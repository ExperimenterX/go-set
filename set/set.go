package set

// Set is a custom data structure to store unique elements.
type Set struct {
    elements map[interface{}]bool
}

// New creates and returns a new Set.
func New() *Set {
    return &Set{elements: make(map[interface{}]bool)}
}

// Add adds an element to the Set.
func (s *Set) Add(value interface{}) {
    s.elements[value] = true
}

// Has checks if an element exists in the Set.
func (s *Set) Has(value interface{}) bool {
    return s.elements[value]
}

// Remove deletes an element from the Set.
func (s *Set) Remove(value interface{}) {
    delete(s.elements, value)
}

// Elements returns all elements in the Set.
func (s *Set) Elements() []interface{} {
    keys := make([]interface{}, 0, len(s.elements))
    for k := range s.elements {
        keys = append(keys, k)
    }
    return keys
}
