package serialization

import (
	"bytes"
	"encoding/gob"

	"github.com/IvanMolodtsov/IoC-Tools/common"
	"github.com/IvanMolodtsov/IoC-Tools/uobject"
)

var DeserializeGob common.Dependency = func(args []common.Any) (common.Any, error) {
	// gob.Register(&uobject.SObject{})
	var (
		mCache = args[0].(*bytes.Buffer)
		data   common.Any
	)
	decCache := gob.NewDecoder(mCache)
	err := decCache.Decode(&data)
	return data, err
}

var SerializeGob common.Dependency = func(args []common.Any) (common.Any, error) {
	// gob.Register(&uobject.SObject{})
	var obj = args[0]
	mCache := new(bytes.Buffer)
	encCache := gob.NewEncoder(mCache)
	err := encCache.Encode(&obj)
	return mCache, err
}

func init() {
	gob.Register(&uobject.SObject{})
}
