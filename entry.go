package registry

import (
  "sync"
)

type Item interface{}

type Entry struct {
  item Item
  cond *sync.Cond
}

func NewEntry(thing interface{}) (ent *Entry) {
  ent = new(Entry)
  ent.cond = sync.NewCond(new(sync.Mutex))
  ent.item = Item(thing)
  return
}

func (e *Entry) NotifyChange() {
  e.cond.Broadcast()
}
