package serialization_test

import (
	"bytes"
	"testing"

	"github.com/IvanMolodtsov/IoC-Tools/common"
	"github.com/IvanMolodtsov/IoC-Tools/ioc"
	"github.com/IvanMolodtsov/IoC-Tools/serialization"
	"github.com/IvanMolodtsov/IoC-Tools/uobject"
)

func TestSObjectJSON(t *testing.T) {
	c, _ := ioc.Resolve[common.ICommand]("IoC.REGISTER", "Deserialize", serialization.DeserializeJSON)
	c.Invoke()
	c, _ = ioc.Resolve[common.ICommand]("IoC.REGISTER", "Serialization.parse", serialization.ParseResult)
	c.Invoke()
	c, _ = ioc.Resolve[common.ICommand]("IoC.REGISTER", "Serialize", serialization.SerializeJSON)
	c.Invoke()
	c, _ = ioc.Resolve[common.ICommand]("IoC.REGISTER", "SObject.new", uobject.NewSObject)
	c.Invoke()
	var jsonString = `{"arr":[{"key":"val"}],"key":"val","obj":{"key":"val"}}`
	result, err := ioc.Resolve[common.UObject]("Deserialize", jsonString)
	if err != nil {
		t.Error(err)
	}
	a, _ := result.Get("arr")
	pa := a.([]common.Any)
	pa[0].(common.UObject).Set("key2", 42.5)

	str, err := ioc.Resolve[string]("Serialize", result)
	if err != nil {
		t.Error(err)
	}
	if string(str) != `{"arr":[{"key":"val","key2":42.5}],"key":"val","obj":{"key":"val"}}` {
		t.Errorf(`expected: {"arr":[{"key":"val","key2":42.5}],"key":"val","obj":{"key":"val"}} got %s`, string(str))
	}
}

func TestSObjectGob(t *testing.T) {
	c, _ := ioc.Resolve[common.ICommand]("IoC.REGISTER", "Deserialize", serialization.DeserializeGob)
	c.Invoke()
	c, _ = ioc.Resolve[common.ICommand]("IoC.REGISTER", "Serialization.parse", serialization.ParseResult)
	c.Invoke()
	c, _ = ioc.Resolve[common.ICommand]("IoC.REGISTER", "Serialize", serialization.SerializeGob)
	c.Invoke()
	c, _ = ioc.Resolve[common.ICommand]("IoC.REGISTER", "SObject.new", uobject.NewSObject)
	c.Invoke()
	dataStruct := make(map[string]common.Any, 1)
	obj, err := ioc.Resolve[common.UObject]("SObject.new", dataStruct)
	obj.Set("key", "value")
	if err != nil {
		t.Error(err)
	}
	buf, err := ioc.Resolve[*bytes.Buffer]("Serialize", obj)

	result, err := ioc.Resolve[common.UObject]("Deserialize", buf)
	if err != nil {
		t.Error(err)
	}

	val, err := result.Get("key")

	if err != nil {
		t.Error(err)
	}
	if val.(string) != "value" {
		t.Errorf("expected 'value' got %s", val)
	}
}
