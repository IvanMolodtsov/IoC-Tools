package serialization

import (
	"encoding/json"

	"github.com/IvanMolodtsov/IoC-Tools/common"
	"github.com/IvanMolodtsov/IoC-Tools/ioc"
)

var DeserializeJSON common.Dependency = func(args []common.Any) (common.Any, error) {
	var jsonString = args[0].(string)
	var result common.Any
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, common.DependencyError{}
	}
	return ioc.Resolve[common.Any]("Serialization.parse", result)
}

var SerializeJSON common.Dependency = func(args []common.Any) (common.Any, error) {
	var obj = args[0]
	value, err := json.Marshal(obj)
	if err != nil {
		return nil, common.DependencyError{}
	}
	return string(value), nil
}
