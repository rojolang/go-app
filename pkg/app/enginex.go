package app

type Engine interface{}

type enginex struct {
	IDs      idGenerator
	Elements elementStore
}
