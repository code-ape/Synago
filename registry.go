package registry

import (
  "sync"
)

type Registry struct {
  sync.RWMutex
  entries map[string]*Entry
  waiting_watchers map[string][]Watcher
}

var GlobalRegistry *Registry


func init() {
  InitGlobalRegistry()
}


func InitGlobalRegistry() {
  GlobalRegistry = new(Registry)
  GlobalRegistry.entries = make(map[string]*Entry)
}

func ItemExists(name string) (exists bool) {
  GlobalRegistry.Lock()
  defer GlobalRegistry.Unlock()
  _,exists = GlobalRegistry[name]
  return
}

func GetItem(name string) (item *Item, err error) {
  GlobalRegistry.RLock()
  defer GlobalRegistry.RUnlock()
  item, err = getItemNoLock(name)
  return
}


func getItemNoLock(name string) (item *Item, err error) {
  var entry *Entry
  entry, err = getEntryNoLock(name)
  item = entry.item
  return
}

func getEntryNoLock(name string) (entry *Entry, err error) {
  var exists bool
  entry,exists = GlobalRegistry[name]
  if !exists {
    err = cust_err.ItemDoesntExist
  }
  return
}

func removeEntryNoLock(name string) (err error) {
  _, err = getItemNoLock(name)
  if err != nil{
    return
  }
  //TODO remove watchers
  delete(m)
}
