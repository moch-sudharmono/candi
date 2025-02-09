package candishared

import (
	"reflect"
	"strconv"
	"strings"
)

type partialUpdateOption struct {
	updateFields map[string]struct{}
	ignoreFields map[string]struct{}
}

// DBUpdateOptionFunc option func
type DBUpdateOptionFunc func(*partialUpdateOption)

// DBUpdateSetUpdatedFields option func
func DBUpdateSetUpdatedFields(fields ...string) DBUpdateOptionFunc {
	return func(o *partialUpdateOption) {
		o.updateFields = make(map[string]struct{})
		for _, field := range fields {
			o.updateFields[field] = struct{}{}
		}
	}
}

// DBUpdateSetIgnoredFields option func
func DBUpdateSetIgnoredFields(fields ...string) DBUpdateOptionFunc {
	return func(o *partialUpdateOption) {
		o.ignoreFields = make(map[string]struct{})
		for _, field := range fields {
			o.ignoreFields[field] = struct{}{}
		}
	}
}

// DBUpdateGORMExtractorKey struct tag key extractor for gorm model
func DBUpdateGORMExtractorKey(structTag reflect.StructTag) string {
	gormTag := structTag.Get("gorm")
	if !strings.HasPrefix(gormTag, "column:") {
		return ""
	}
	return strings.Split(strings.TrimPrefix(gormTag, "column:"), ";")[0]
}

// DBUpdateMongoExtractorKey struct tag key extractor for mongo model
func DBUpdateMongoExtractorKey(structTag reflect.StructTag) string {
	return strings.TrimSuffix(structTag.Get("bson"), ",omitempty")
}

// DBUpdateTools for construct selected field to update
type DBUpdateTools struct {
	KeyExtractorFunc func(structTag reflect.StructTag) string
	IgnoredFields    []string
}

// ToMap method
func (d DBUpdateTools) ToMap(data interface{}, opts ...DBUpdateOptionFunc) map[string]interface{} {
	var (
		o            partialUpdateOption
		updateFields = make(map[string]interface{}, 0)
	)

	for _, opt := range opts {
		opt(&o)
	}

	dataValue := reflect.ValueOf(data)
	dataType := reflect.TypeOf(data)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
		dataType = dataType.Elem()
	}
	isPartial := len(o.updateFields) > 0 || len(o.ignoreFields) > 0

	for i := 0; i < dataValue.NumField(); i++ {
		fieldValue := dataValue.Field(i)
		fieldType := dataType.Field(i)

		if fieldType.Anonymous {
			for k, v := range d.ToMap(fieldValue.Interface(), opts...) {
				updateFields[k] = v
			}
			continue
		}

		isIgnore, _ := strconv.ParseBool(fieldType.Tag.Get("ignoreUpdate"))
		key := strings.TrimSuffix(fieldType.Tag.Get("json"), ",omitempty")
		if d.KeyExtractorFunc != nil {
			key = d.KeyExtractorFunc(fieldType.Tag)
		}
		if key == "" || key == "-" || isIgnore {
			continue
		}

		val := fieldValue.Interface()
		if fieldValue.Kind() == reflect.Pointer {
			val = fieldValue.Elem().Interface()
		}

		if !isPartial {
			updateFields[key] = val
			continue
		}

		_, isFieldUpdated := o.updateFields[fieldType.Name]
		_, isFieldIgnored := o.ignoreFields[fieldType.Name]
		if (isFieldUpdated && len(o.updateFields) > 0) || (!isFieldIgnored && len(o.ignoreFields) > 0) {
			updateFields[key] = val
		}
	}

	for _, ignored := range d.IgnoredFields {
		delete(updateFields, ignored)
	}
	return updateFields
}
