package uobject

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"sync"

	"github.com/IvanMolodtsov/IoC-Tools/common"

	"github.com/IvanMolodtsov/IoC-Tools/ioc"
)

type SObject struct {
	data map[string]common.Any
	lock sync.RWMutex
}

func (o *SObject) Get(key string) (common.Any, error) {
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

func (o *SObject) Set(key string, value common.Any) error {
	o.lock.Lock()
	o.data[key] = value
	o.lock.Unlock()
	return nil
}

func (o *SObject) Remove(key string) {
	o.lock.Lock()
	delete(o.data, key)
	o.lock.Unlock()
}

func (o *SObject) MarshalJSON() ([]byte, error) {
	o.lock.Lock()
	defer o.lock.Unlock()
	return json.Marshal(o.data)
}

func (o *SObject) UnmarshalJSON(b []byte) error {
	var (
		err  error
		data map[string]common.Any
	)
	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	} else {
		for k, val := range data {
			result, err := ioc.Resolve[common.Any]("Serialization.parse", val)
			if err != nil {
				return err
			}
			o.Set(k, result)
		}
		return nil
	}
}

func (o *SObject) GobEncode() ([]byte, error) {
	var buffer = new(bytes.Buffer)
	var encoder = gob.NewEncoder(buffer)
	o.lock.Lock()
	err := encoder.Encode(o.data)
	o.lock.Unlock()
	return buffer.Bytes(), err
}

func (o *SObject) GobDecode(data []byte) error {
	var (
		err     error
		decoded map[string]common.Any
		buffer  = bytes.NewBuffer(data)
	)
	if o.data == nil {
		o.data = make(map[string]common.Any)
	}

	var decoder = gob.NewDecoder(buffer)
	err = decoder.Decode(&decoded)
	if err != nil {
		return err
	} else {
		for k, val := range decoded {
			result, err := ioc.Resolve[common.Any]("Serialization.parse", val)
			if err != nil {
				return err
			}
			o.Set(k, result)
		}
		return nil
	}
}

var NewSObject common.Dependency = func(a []common.Any) (common.Any, error) {
	var o = new(SObject)
	if len(a) == 1 {
		o.data = a[0].(map[string]common.Any)
	} else {
		o.data = map[string]common.Any{}
	}
	return o, nil
}
