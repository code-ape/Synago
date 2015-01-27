package registry

func RegisterItem(name string, thing interface{}) (err error) {
  GlobalRegistry.Lock()
  defer GlobalRegistry.Unlock()

  if ItemExists(name) {
    err = cust_err.ItemNameRegistered
    return
  }
  updateItemNoLock(name, thing)
  return
}


func DeregisteredItem(name string) error {
  GlobalRegistry.Lock()
  defer GlobalRegistry.Unlock()
  return removeItemNoLock(name)
}
