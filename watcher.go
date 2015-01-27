package registry

import (
  "sync"
)

type WatchArg uint

const (
  WaitForCreation WatcherArg = iota
  WaitAfterDeletion
  UpdateOnDeletionOnly
  UpdateOnCreationOnly
)

type Watcher struct {
  key string
  entry *Entry
  cond *sync.Cond
}

func NewWatcher(key string) (watcher *Watcher) {
  watcher := new(Watcher)
  watcher.key = key
  return
}

func WatchKey(key string) (watcher *Watcher) {
  watcher = NewWatcher(key)
  return
}

func WatchKeyIfExists(key string) *Watcher {

}


func (w *Watcher) WaitForUpdate() {
  w.cond.L.Lock()
  w.cond.Wait()
  w.cond.L.Unlock()
}




func main() {
  w := r.NewWatcher()
  err := w.SetKey("foo").SetArgs(WaitForCreation).Commit()
  if err != nil {
    fmt.Println(err)
    panic
  }
  go func() {
    w.WaitForUpdate()
    fmt.Println("Update!")
  }()
}
