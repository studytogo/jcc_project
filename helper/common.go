package helper

import (
	"github.com/astaxie/beego/logs"
	"reflect"
	"strconv"
)

// 要求srcStruct中的每一个都是string类型的
func ReflectiveStruct(newStrcut interface{}, srcStruct interface{}) error {
	v := reflect.ValueOf(newStrcut).Elem()
	sv := reflect.ValueOf(srcStruct).Elem()
	for i := 0; i < v.NumField(); i++ {
		sf := v.Type().Field(i)
		// 获取原结构体的value
		value := sv.FieldByName(sf.Name)
		if !value.IsValid() || value.String() == "" {
			continue
		}

		// 填充对应新结构体value
		newValue := v.FieldByName(sf.Name)
		switch sf.Type.Name() {
		case "string":
			newValue.SetString(value.String())
		case "int", "int8", "int16", "int32", "int64":
			int64Value, err := strconv.ParseInt(value.String(), 10, 64)
			if err != nil {
				logs.Error("strconv.ParseInt", "-", sf.Name, ":", value.String())
				return err
			}
			newValue.SetInt(int64Value)
		case "uint", "uint8", "uint16", "uint32", "uint64":
			uint64Value, err := strconv.ParseUint(value.String(), 10, 64)
			if err != nil {
				logs.Error("strconv.ParseUint", "-", sf.Name, ":", value.String())
				return err
			}
			newValue.SetUint(uint64Value)
		case "float32", "float64":
			float64Value, err := strconv.ParseFloat(value.String(), 64)
			if err != nil {
				logs.Error("strconv.ParseFloat", "-", sf.Name, ":", value.String())
				return err
			}
			newValue.SetFloat(float64Value)
		default:
			//v.FieldByName(sf.Name).Set(sv.FieldByName(sf.Name))
		}
	}
	return nil
}
