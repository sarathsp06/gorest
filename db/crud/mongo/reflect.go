package mongo

import (
	"fmt"
	"reflect"
)

// StructField represents a struct field with field type and value
type StructField struct {
	Type  reflect.StructField
	Value reflect.Value
}

// StructToMap converts a struct to map of string to interface
// Its uses the tag 'map' or  given `tag` to check the name to be used
// if the value is zero value and withValue is true no key is added
// if there are embeded structs they are also traversed recursively
// in case of embeded structs having same field name and struct tag the last field in order will be considered
func StructToMap(in interface{}, tag string, withValue bool) (m map[string]interface{}, err error) {
	if tag == "" {
		tag = "map"
	}
	m = make(map[string]interface{})

	fields, err := flattenStruct(in)
	if err != nil {
		return nil, err
	}
	for _, field := range fields {
		//if embededstruct deal with it
		if field.Type.Anonymous && field.Type.Tag.Get(tag) == "" {
			embededVals, _ := StructToMap(field.Value.Interface(), tag, withValue)
			for key, value := range embededVals {
				m[key] = value
			}
			continue
		}
		if withValue == false || !isEmptyValue(field.Value) {
			if tagv := field.Type.Tag.Get(tag); tagv != "" {
				m[tagv] = field.Value.Interface()
			} else {
				m[field.Type.Name] = field.Value.Interface()
			}
		}
	}
	return
}

// StructToAddrMap returns a map of fielname(or tag) to field variable reference
// if not struct or pointer to struct it fails
func StructToAddrMap(in interface{}, tag string) (map[string]reflect.Value, error) {
	if tag == "" {
		tag = "map"
	}
	m := make(map[string]reflect.Value)
	fields, err := flattenStruct(in)
	if err != nil {
		return nil, err
	}
	for _, field := range fields {
		//if embededstruct deal with it
		if field.Type.Anonymous && field.Type.Tag.Get(tag) == "" {
			embededVals, _ := StructToAddrMap(field.Value.Addr().Interface(), tag)
			for key, value := range embededVals {
				m[key] = value
			}
			continue
		}
		if tagv := field.Type.Tag.Get(tag); tagv != "" {
			m[tagv] = field.Value
		} else {
			m[field.Type.Name] = field.Value
		}

	}
	return m, nil
}

// flattenStruct converts  a struct to array of fields
func flattenStruct(in interface{}) ([]StructField, error) {
	var structFields []StructField
	v := reflect.ValueOf(in)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("only accepts structs; got %T", v)
	}
	tp := v.Type()
	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		value := v.Field(i)
		structFields = append(structFields, StructField{Type: field, Value: value})
	}
	return structFields, nil
}

//MapToStruct converts map to given struct
func MapToStruct(m map[string]interface{}, out interface{}) error {
	addressMap, err := StructToAddrMap(out, "")
	if err != nil {
		return err
	}
	for key, value := range m {
		addressMap[key].Set(reflect.ValueOf(value))
	}
	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
		return true
	}
	return false
}
