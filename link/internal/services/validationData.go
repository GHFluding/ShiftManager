package services

import (
	"fmt"
	"strconv"
)

func getInt64(data map[string]interface{}, key string) (int64, error) {
	val, exists := data[key]
	if !exists {
		return 0, fmt.Errorf("missing required field: %s", key)
	}

	switch v := val.(type) {
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, fmt.Errorf("invalid type for field %s: %T", key, val)
	}
}

func getString(data map[string]interface{}, key string) (string, error) {
	val, exists := data[key]
	if !exists {
		return "", fmt.Errorf("missing required field: %s", key)
	}

	s, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("field %s is not a string: %T", key, val)
	}
	return s, nil
}
func getBool(val interface{}) (bool, error) {
	switch v := val.(type) {
	case bool:
		return v, nil
	case string:
		return v == "true", nil
	case int, float64:
		return v != 0, nil
	default:
		return false, fmt.Errorf("invalid bool type: %T", val)
	}
}
func getInt64Optional(val interface{}) (int64, error) {
	switch v := val.(type) {
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, fmt.Errorf("invalid type for int64: %T", val)
	}
}
