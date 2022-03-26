package uobject

import (
	"sync"

	"github.com/IvanMolodtsov/IoC-Tools/common"
)

type BasicObject struct {
	data map[string]common.Any
	lock sync.RWMutex
}

func (o *BasicObject) Get(key string) (common.Any, error) {
	var value common.Any
	var err error
	o.lock.RLock()
	value, ok := o.data[key]
	o.lock.RUnlock()
	if !ok {
		err = common.GetterError{}
	}
	return value, err
}

func (o *BasicObject) Set(key string, value common.Any) error {
	o.lock.Lock()
	o.data[key] = value
	o.lock.Unlock()
	return nil
}

func (o *BasicObject) Remove(key string) {
	o.lock.Lock()
	delete(o.data, key)
	o.lock.Unlock()
}
