package service

import (
	"errors"
	"fmt"
	"go-hex-temp/internal/core/domain"
	"go-hex-temp/internal/core/schema"
	"strconv"
	"strings"
)

const maxLimit = 10000

var allowedOps map[string]bool

type QCompiler struct {
	schemaInfo schema.SchemaInfo
}

func NewQCompiler() *QCompiler {
	initOps()
	sc := make(schema.SchemaInfo)
	sc.Register(domain.User{}) // Register all needed domain structs here
	return &QCompiler{sc}
}

func initOps() {
	allowedOps = make(map[string]bool)

	// helper to register allowed ops for a type
	registerOps := func(t string, ops ...domain.Operator) {
		for _, op := range ops {
			key := t + ":" + string(op)
			allowedOps[key] = true
		}
	}

	registerOps("string", domain.Eq, domain.Ne, domain.Like, domain.Regex, domain.In, domain.Nin, domain.Exists, domain.Null)
	registerOps("int", domain.Eq, domain.Ne, domain.Gt, domain.Lt, domain.Gte, domain.Lte, domain.Between, domain.In, domain.Nin, domain.Exists, domain.Null)
	registerOps("uint", domain.Eq, domain.Ne, domain.Gt, domain.Lt, domain.Gte, domain.Lte, domain.Between, domain.In, domain.Nin, domain.Exists, domain.Null)
	registerOps("bool", domain.Eq, domain.Ne, domain.Exists, domain.Null)
	registerOps("float64", domain.Eq, domain.Ne, domain.Gt, domain.Lt, domain.Gte, domain.Lte, domain.Between, domain.In, domain.Nin, domain.Exists, domain.Null)
}

func isAllowedOp(fieldType string, op domain.Operator) bool {
	key := fieldType + ":" + string(op)
	return allowedOps[key]
}

// compile validates the Query against the schema
func (c *QCompiler) Compile(query *domain.Query, dom schema.Schema) (*domain.Query, error) {
	scName := dom.ScName()

	// Validate Filters
	for field, conditions := range query.Filter {
		fieldType := strings.ToLower(c.schemaInfo.GetType(scName, field))
		if fieldType == "" {
			return nil, fmt.Errorf("unknown field: %s in schema %s", field, scName)
		}

		for op, values := range conditions {
			if !isAllowedOp(fieldType, op) {
				return nil, fmt.Errorf("operator %s not allowed on field %s of type %s", op, field, fieldType)
			}
			// type conversion
			for i, val := range values {
				castedVal, err := castValue(fieldType, val)
				if err != nil {
					return nil, fmt.Errorf("invalid value for field %s operator %s: %v", field, op, err)
				}
				values[i] = castedVal
			}

		}
	}

	// Validate Pagination if exists
	if query.Pagination != nil {

		if query.Pagination.Limit <= 0 {
			query.Pagination.Limit = 10
		}
		if query.Pagination.Limit > maxLimit {
			return nil, fmt.Errorf("pagination limit must not exceed %d", maxLimit)
		}

		// Validate Sort fields and directions
		for _, sort := range query.Pagination.Sorts {
			fieldType := strings.ToLower(c.schemaInfo.GetType(scName, sort.Key))
			if fieldType == "" {
				return nil, fmt.Errorf("unknown sort field: %s in schema %s", sort.Key, scName)
			}

			dir := sort.Direction
			if dir != domain.Asc && dir != domain.Desc {
				return nil, fmt.Errorf("invalid sort direction: %s for field %s", sort.Direction, sort.Key)
			}
		}
	}

	return query, nil
}

func castValue(fieldType string, val any) (any, error) {
	switch fieldType {
	case "string":
		// if val is string, just assert; else try fmt.Sprintf to convert
		if s, ok := val.(string); ok {
			return s, nil
		}
		return fmt.Sprintf("%v", val), nil

	case "int":
		return toInt(val)

	case "uint":
		return toUint(val)

	case "bool":
		return toBool(val)

	case "float64":
		return toFloat64(val)

	default:
		return nil, fmt.Errorf("unsupported field type %s", fieldType)
	}
}

func toInt(val any) (int, error) {
	switch v := val.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, errors.New("cannot convert to int")
	}
}

func toUint(val any) (uint, error) {
	errNegative := errors.New("negative value for uint")
	switch v := val.(type) {
	case uint:
		return v, nil
	case uint8:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint64:
		return uint(v), nil
	case int:
		if v < 0 {
			return 0, errNegative
		}
		return uint(v), nil
	case float32:
		if v < 0 {
			return 0, errNegative
		}
		return uint(v), nil
	case float64:
		if v < 0 {
			return 0, errNegative
		}
		return uint(v), nil
	case string:
		ui64, err := strconv.ParseUint(v, 10, 64)
		return uint(ui64), err
	default:
		return 0, errors.New("cannot convert to uint")
	}
}

func toBool(val any) (bool, error) {
	switch v := val.(type) {
	case bool:
		return v, nil
	case string:
		return strconv.ParseBool(v)
	default:
		return false, errors.New("cannot convert to bool")
	}
}

func toFloat64(val any) (float64, error) {
	switch v := val.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, errors.New("cannot convert to float64")
	}
}
