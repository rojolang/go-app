package app

import "github.com/maxence-charriere/go-app/v9/pkg/errors"

type Engine interface{}

type enginex struct {
	IDs      idGenerator
	Elements elementStore

	root HTMLBody
}

// newEngineX creates a new engine instance using the provided body element as
// its root.
// It sets up the descriptor for the root element and attempts to mount it.
// The function will panic if mounting the root element fails.
func newEngineX(root HTMLBody) *enginex {
	e := &enginex{}

	root.setDescriptor(elementDescriptor{
		ID:        e.IDs.NewID(),
		Depth:     1,
		JSElement: Window().Get("document").Get("body"),
	})

	if err := e.Elements.Mount(root); err != nil {
		panic(errors.New("mounting root element failed").Wrap(err))
	}

	e.root = root
	return e
}
