package app

// A custom type used for representing identifiers.
type ID uint

// An id generator that manages the generation and recycling of unique IDs.
type idGenerator struct {
	nextID      ID
	reusableIDs map[ID]struct{}
}

// NewID returns a new unique ID. If any IDs have been marked as reusable,
// one of them is returned. Otherwise, a newly incremented ID is returned.
func (g *idGenerator) NewID() ID {
	if len(g.reusableIDs) != 0 {
		for id := range g.reusableIDs {
			delete(g.reusableIDs, id)
			return id
		}
	}

	g.nextID++
	return g.nextID
}

// ReuseID marks an ID as reusable. This ID may be returned by future calls
// to NewID. The function is a no-op if the given ID is zero.
func (g *idGenerator) ReuseID(v ID) {
	if v == 0 {
		return
	}

	if g.reusableIDs == nil {
		g.reusableIDs = make(map[ID]struct{})
	}
	g.reusableIDs[v] = struct{}{}
}
