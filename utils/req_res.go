package utils

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
)

func NewRequestBody[T any](proxyRequest events.APIGatewayProxyRequest, additionalData T) (T, error) {
	if proxyRequest.Body == "" {
		return additionalData, nil
	}

	var result T
	if err := json.Unmarshal([]byte(proxyRequest.Body), &result); err != nil {
		return additionalData, fmt.Errorf("failed to unmarshal request body: %v", err)
	}
	if err := validateKeys(additionalData, result); err != nil {
		return additionalData, err
	}

	return result, nil
}

func validateKeys(additionalData, result interface{}) error {
	additionalDataValue := reflect.ValueOf(additionalData)
	resultValue := reflect.ValueOf(result)

	if additionalDataValue.Kind() != reflect.Struct || resultValue.Kind() != reflect.Struct {
		return fmt.Errorf("both additionalData and result must be structs")
	}

	additionalDataType := additionalDataValue.Type()
	resultType := resultValue.Type()

	for i := 0; i < additionalDataType.NumField(); i++ {
		additionalField := additionalDataType.Field(i)
		resultField, exists := resultType.FieldByName(additionalField.Name)

		if !exists {
			return fmt.Errorf("field %s exists in additionalData but not in result", additionalField.Name)
		}

		if additionalField.Type != resultField.Type {
			return fmt.Errorf("field %s has different types in additionalData and result", additionalField.Name)
		}

		// フィールドの値が空でないことを確認
		resultFieldValue := resultValue.FieldByName(additionalField.Name)
		if isEmptyValue(resultFieldValue) {
			return fmt.Errorf("field %s in result is empty", additionalField.Name)
		}
	}

	// resultに余分なフィールドがないか確認
	for i := 0; i < resultType.NumField(); i++ {
		resultField := resultType.Field(i)
		if _, exists := additionalDataType.FieldByName(resultField.Name); !exists {
			return fmt.Errorf("field %s exists in result but not in additionalData", resultField.Name)
		}
	}

	return nil
}

// isEmptyValue checks if a reflect.Value is its zero value.
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
	case reflect.Struct:
		// 構造体の場合、すべてのフィールドが空かどうかを再帰的にチェック
		for i := 0; i < v.NumField(); i++ {
			if !isEmptyValue(v.Field(i)) {
				return false
			}
		}
		return true
	}
	return false
}

func BodyToJSON[T any](body T) (string, error) {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	return string(bodyJson), nil
}

func JSONtoBody[T any](jsonStr string, body *T) error {
	return json.Unmarshal([]byte(jsonStr), body)
}
