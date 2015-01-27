package registry


func UpdateItem(name string, thing interface{}) (err error) {
  GlobalRegistry.Lock()
  defer GlobalRegistry.Unlock()

  if !ItemExists(name) {
    err = cust_err.ItemDoesntExist
    return
  }
  updateItemNoLock(name, thing)
  return
}

func UpdateItemHARD(name string, thing interface{}) {
  GlobalRegistry.Lock()
  defer GlobalRegistry.Unlock()
  updateItemNoLock(name, thing)
}

func updateItemNoLock(name string, thing interface{}) {
  GlobalRegistry[name] = NewEntry(thing)
}
