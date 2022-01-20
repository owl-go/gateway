package utils

import (
	"encoding/json"
	"reflect"
	"strconv"
)

func InterfaceToInt(infVal interface{}) (rValue int) {
	switch infVal.(type) {
	case nil:
		rValue = 0
	case int:
		rValue = infVal.(int)
	case int32:
		rValue = int(infVal.(int32))
	case int64:
		rValue = int(infVal.(int64))
	case float32:
		rValue = int(infVal.(float32))
	case float64:
		rValue = int(infVal.(float64))
	case []byte:
		zValue := string(infVal.([]byte))
		if iValue, err := strconv.ParseInt(zValue, 10, 64); err == nil {
			rValue = int(iValue)
		}
	case string:
		if iValue, err := strconv.ParseInt(infVal.(string), 10, 64); err == nil {
			rValue = int(iValue)
		}
	}
	return
}

func InterfaceToInt32(infVal interface{}) (rValue int32) {
	switch infVal.(type) {
	case nil:
		rValue = 0
	case int:
		rValue = int32(infVal.(int))
	case int32:
		rValue = int32(infVal.(int32))
	case int64:
		rValue = int32(infVal.(int64))
	case float32:
		rValue = int32(infVal.(float32))
	case float64:
		rValue = int32(infVal.(float64))
	case []byte:
		zValue := string(infVal.([]byte))
		if iValue, err := strconv.ParseInt(zValue, 10, 64); err == nil {
			rValue = int32(iValue)
		}
	case string:
		if iValue, err := strconv.ParseInt(infVal.(string), 10, 64); err == nil {
			rValue = int32(iValue)
		}
	}
	return
}

func InterfaceToInt64(infVal interface{}) (rValue int64) {
	switch infVal.(type) {
	case nil:
		rValue = 0
	case int:
		rValue = int64(infVal.(int))
	case int32:
		rValue = int64(infVal.(int32))
	case int64:
		rValue = int64(infVal.(int64))
	case float32:
		rValue = int64(infVal.(float32))
	case float64:
		rValue = int64(infVal.(float64))
	case []byte:
		zValue := string(infVal.([]byte))
		if iValue, err := strconv.ParseInt(zValue, 10, 64); err == nil {
			rValue = int64(iValue)
		}
	case string:
		if iValue, err := strconv.ParseInt(infVal.(string), 10, 64); err == nil {
			rValue = int64(iValue)
		}
	}
	return
}

func InterfaceToString(infVal interface{}) (rValue string) {
	switch infVal.(type) {
	case nil:
		rValue = ""
	case int:
		rValue = strconv.FormatInt(int64(infVal.(int)), 10)
	case int32:
		rValue = strconv.FormatInt(int64(infVal.(int32)), 10)
	case int64:
		rValue = strconv.FormatInt(infVal.(int64), 10)
	case float32:
		rValue = strconv.FormatInt(int64(infVal.(float32)), 10)
	case float64:
		rValue = strconv.FormatInt(int64(infVal.(float64)), 10)
	case []byte:
		rValue = string(infVal.([]byte))
	case string:
		rValue = infVal.(string)
	default:
		if bytes, err := json.Marshal(infVal); err == nil {
			rValue = string(bytes)
		}
	}
	return
}

func InterfaceToBool(infVal interface{}) (rValue bool) {
	switch infVal.(type) {
	case nil:
		rValue = false
	case bool:
		rValue = infVal.(bool)
	case int:
		rValue = infVal.(int) != 0
	case int32:
		rValue = infVal.(int32) != 0
	case int64:
		rValue = infVal.(int64) != 0
	case float32:
		rValue = int32(infVal.(float32)) != 0
	case float64:
		rValue = int32(infVal.(float64)) != 0
	case []byte:
		zValue := string(infVal.([]byte))
		if iValue, err := strconv.ParseInt(zValue, 10, 64); err == nil {
			rValue = int32(iValue) != 0
		}
	case string:
		if iValue, err := strconv.ParseInt(infVal.(string), 10, 64); err == nil {
			rValue = int32(iValue) != 0
		}
	}
	return
}

func InterfaceToJsonString(infVal interface{}) (string, bool) {
	infBytes, err := json.Marshal(infVal)
	if err != nil {
		return "", false
	}
	return string(infBytes[:]), true
}

func InterfaceToStringArray(infVale interface{}) []string {
	str := make([]string, 0)
	iArr, ok := infVale.([]string)
	if ok {
		return iArr
	} else {
		iArr, ok := infVale.([]interface{})
		if ok {
			for _, v := range iArr {
				if _, ok := v.(string); ok {
					str = append(str, v.(string))
				}
			}
		}
		return str
	}
	return iArr
}

func StringToKindInterface(infKind reflect.Kind, zVal string) (rValue interface{}) {
	switch infKind {
	case reflect.Bool:
		rValue, _ = strconv.ParseBool(zVal)
	case reflect.Int:
		rValue, _ = strconv.Atoi(zVal)
	case reflect.Int32:
		iValue, _ := strconv.Atoi(zVal)
		rValue = int32(iValue)
	case reflect.Int64:
		iValue, _ := strconv.Atoi(zVal)
		rValue = int64(iValue)
	case reflect.Float32:
		fValue, _ := strconv.ParseFloat(zVal, 64)
		rValue = int64(fValue)
	case reflect.Float64:
		fValue, _ := strconv.ParseFloat(zVal, 64)
		rValue = int64(fValue)
	case reflect.String:
		rValue = zVal
	case reflect.Slice:
		var vInf []interface{}
		json.Unmarshal([]byte(zVal), &vInf)
		rValue = vInf
	}
	return
}
