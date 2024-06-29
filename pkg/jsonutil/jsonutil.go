package jsonutil

import (
	"encoding/json"
	"reflect"
)

type VariableMap map[string]interface{}

func (r VariableMap) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for key, val := range r {
		if val == nil {
			m[key] = nil
			continue
		}
		v := reflect.Indirect(reflect.ValueOf(val))
		dataKind := v.Kind()
		dataType := v.Type()
		switch {
		case dataType.PkgPath() == "encoding/json" && dataType.Name() == "Number":
			jn := val.(json.Number)
			i, err := jn.Float64()
			if err != nil {
				return nil, err
			}
			m[key] = i
		case dataKind == reflect.Map:
			m[key] = VariableMap(val.(map[string]interface{}))
		default:
			m[key] = val
		}
	}
	return json.Marshal(m)
}
