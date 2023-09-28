package app

import (
	"reflect"

	"github.com/maxence-charriere/go-app/v9/pkg/errors"
)

type Engine interface{}

type engineX struct {
	ids      idGenerator
	elements elementStore
	root     HTMLBody
	page     Page
}

// newEngineX creates a new engine instance using the provided body element as
// its root.
// It sets up the descriptor for the root element and attempts to mount it.
// The function will panic if mounting the root element fails.
func newEngineX(root HTMLBody) *engineX {
	e := &engineX{}

	root.setDescriptor(elementDescriptor{
		ID:        e.ids.NewID(),
		Depth:     1,
		JSElement: Window().Get("document").Get("body"),
	})

	if err := e.elements.Mount(root); err != nil {
		panic(errors.New("mounting root element failed").Wrap(err))
	}

	e.root = root
	return e
}

func (e *engineX) load(page Page, compo Composer) error {
	if page.URL() == nil {
		return errors.New("page url is missing")
	}
	e.page = page

	if children := e.root.getChildren(); len(children) != 0 {
		rootCompo := children[0]
		if reflect.TypeOf(rootCompo) == reflect.TypeOf(compo) {
			return e.update(rootCompo, compo)
		}

	}

}

func (e *engineX) mount(depth uint, v UI) error {
	panic("not implemented")
}

func (e *engineX) dismount(v UI) {
	panic("not implemented")
}

func (e *engineX) update(a, b UI) error {
	panic("not implemented")
}
