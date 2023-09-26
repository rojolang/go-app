package app

import (
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
)

// elementDescriptor holds metadata about a UI element, such as its ID, depth,
// and associated JavaScript value.
type elementDescriptor struct {
	ID        ID
	Depth     uint
	JSElement Value
}

// JSValue returns the JavaScript value associated with the element descriptor.
func (e elementDescriptor) JSValue() Value {
	return e.JSElement
}

// elementStore maintains a collection of UI elements identified by their IDs.
type elementStore struct {
	elements map[ID]UI
}

// Mount attempts to mount a UI element into the store. It returns an error if
// the element's descriptor has invalid fields or if the element is already
// mounted.
func (s *elementStore) Mount(v UI) error {
	descriptor := v.descriptor()
	if descriptor.ID == 0 || descriptor.Depth == 0 {
		return errors.New("descriptor id or depth is not set").
			WithTag("id", descriptor.ID).
			WithTag("depth", descriptor.Depth)
	}
	if _, ok := s.elements[descriptor.ID]; ok {
		return errors.New("ui element is already mounted").
			WithTag("id", descriptor.ID).
			WithTag("depth", descriptor.Depth)
	}

	if s.elements == nil {
		s.elements = make(map[ID]UI)
	}
	s.elements[descriptor.ID] = v
	return nil
}

// Mounted checks if a given UI element is currently mounted in the store.
// It returns true if mounted, otherwise false.
func (s *elementStore) Mounted(v UI) bool {
	_, mounted := s.elements[v.descriptor().ID]
	return mounted
}

// Dismount removes a UI element from the store.
func (s *elementStore) Dismount(v UI) {
	delete(s.elements, v.descriptor().ID)
}

// Update attempts to update a UI element in the store. It returns an error if
// the element is not already mounted or if it mismatch the current UI element.
func (s *elementStore) Update(v UI) error {
	descriptor := v.descriptor()
	e, mounted := s.elements[descriptor.ID]
	if !mounted {
		return errors.New("ui element is not mounted").
			WithTag("id", descriptor.ID).
			WithTag("depth", descriptor.Depth)
	}
	if currentDescriptor := e.descriptor(); currentDescriptor.Depth != descriptor.Depth {
		return errors.New("updated ui element does not match current ui element").
			WithTag("id", descriptor.ID).
			WithTag("current-depth", currentDescriptor.Depth).
			WithTag("new-depth", descriptor.Depth)
	}

	s.elements[descriptor.ID] = v
	return nil
}
