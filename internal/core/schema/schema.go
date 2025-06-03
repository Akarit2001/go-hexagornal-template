package schema

import (
	"reflect"
	"strings"
)

type Schema interface {
	ScName() string
}

// FieldType maps field names to their type names.
type fieldType map[string]string

// SchemaInfo maps schema names to their field type definitions.
type SchemaInfo map[string]fieldType

// Register uses reflection to extract field names and types.
func (sf *SchemaInfo) Register(dom Schema) {
	if *sf == nil {
		*sf = make(SchemaInfo)
	}

	t := reflect.TypeOf(dom)
	// If it's a pointer, get the underlying type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fieldMap := make(fieldType)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldMap[strings.ToLower(field.Name)] = field.Type.Name()
	}
	// use reflect name

	(*sf)[t.Name()] = fieldMap
}

// GetType returns the type name of a field in a schema.
func (sf SchemaInfo) GetType(scName, scFields string) string {
	if fields, ok := sf[scName]; ok {
		return fields[scFields]
	}
	return ""
}
