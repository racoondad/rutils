package mapstructure

import (
	"reflect"
)

func StructToMap(inStructure interface{}) (mp map[string]interface{}) {
	mp = make(map[string]interface{})

	KType := reflect.TypeOf(inStructure)
	KValue := reflect.ValueOf(inStructure)

	for i := 0; i < KType.NumField(); i++ {
		if KType.Field(i).Tag.Get("mapstructure") != "" {
			mp[KType.Field(i).Tag.Get("mapstructure")] = KValue.Field(i).Interface()
		} else {
			mp[KType.Field(i).Name] = KValue.Field(i).Interface()
		}
	}
	return
}
