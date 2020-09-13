package web_net

import (
	"fmt"
	"reflect"
	"sort"
)

type Schema struct {
	Name   *string  `json:"name,omitempty"`
	Type   string   `json:"type"`
	Fields []Schema `json:"fields,omitempty"`
}

func (s *Schema) concatenatedFieldNames() string {
	fn := ""

	for _, f := range s.Fields {
		if f.Name != nil {
			fn += *f.Name
		}
	}

	return fn
}

func GenerateSchema(in interface{}) *Schema {
	v := reflect.Indirect(reflect.ValueOf(in))
	return processItem(v)
}

func processItem(v reflect.Value) *Schema {
	var s Schema

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		s.Type = "integer"
	case reflect.String:
		s.Type = "string"
	case reflect.Float32, reflect.Float64:
		s.Type = "float"
	case reflect.Bool:
		s.Type = "bool"
	case reflect.Slice:
		s.Type = "array"
		s.Fields = processSlice(v).Fields
	case reflect.Struct:
		s.Type = "object"
		s.Fields = processStruct(v).Fields
	case reflect.Interface:
		return processItem(v.Elem())
	case reflect.Ptr:
		return processItem(reflect.Indirect(reflect.New(v.Type().Elem())))
	case reflect.Map:
		s.Type = "object"
		s.Fields = processMap(v).Fields
	default:
		fmt.Println("Unknown type", v.Kind().String())
	}

	return &s
}

func processSlice(v reflect.Value) *Schema {
	var fields []Schema

	fieldTypeMap := map[string]bool{}

	for i := 0; i < v.Len(); i++ {
		s := processItem(reflect.Indirect(v.Index(i)))

		if _, ok := fieldTypeMap[s.concatenatedFieldNames()]; !ok {
			fields = append(fields, *s)
			fieldTypeMap[s.concatenatedFieldNames()] = true
		}
	}

	return &Schema{
		Type:   "array",
		Fields: fields,
	}
}

func processMap(v reflect.Value) *Schema {
	var fields []Schema

	iter := v.MapRange()

	for iter.Next() {
		s := processItem(iter.Value())
		s.Name = strptr(iter.Key().String())
		fields = append(fields, *s)
	}

	sort.SliceStable(fields, func(i, j int) bool {
		a := fields[i]
		b := fields[j]

		return *a.Name < *b.Name
	})

	return &Schema{
		Fields: fields,
	}
}

func processStruct(v reflect.Value) *Schema {
	var fields []Schema

	for i := 0; i < v.NumField(); i++ {
		s := processItem(v.Field(i))

		str := v.Type().Field(i).Name
		s.Name = &str

		fields = append(fields, *s)
	}

	return &Schema{
		Type:   "object",
		Fields: fields,
	}
}

func strptr(s string) *string {
	return &s
}
