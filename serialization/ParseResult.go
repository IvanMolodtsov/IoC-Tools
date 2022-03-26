package serialization

import (
	"github.com/IvanMolodtsov/IoC-Tools/common"
	"github.com/IvanMolodtsov/IoC-Tools/ioc"
)

var ParseResult common.Dependency = func(args []common.Any) (common.Any, error) {
	var (
		err  error
		data = args[0]
	)
	switch data.(type) {
	default:
		{
			return nil, common.DependencyError{}
		}
	case bool, float64, string, nil:
		{
			return data, nil
		}
	case []interface{}:
		{
			var arr = make([]common.Any, len(data.([]interface{})))
			for k, val := range data.([]interface{}) {
				arr[k], err = ioc.Resolve[common.Any]("Serialization.parse", val)
				if err != nil {
					return nil, err
				}
			}
			return arr, nil
		}
	case map[string]interface{}:
		{
			obj, err := ioc.Resolve[common.UObject]("SObject.new")
			if err != nil {
				return nil, err
			}
			for k, val := range data.(map[string]interface{}) {
				result, err := ioc.Resolve[common.Any]("Serialization.parse", val)
				if err != nil {
					return nil, err
				}
				obj.Set(k, result)
			}
			return obj, nil
		}
	}
}
