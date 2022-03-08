package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"
)

//Empty 类似于 PHP 的 empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
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
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

//把JSON 数据转化为 map
func JSONToMap(filePath string) map[string]string {
	//打开读取json文件
	jsonFile, err := os.Open(filePath)
	//结束的时候关闭文件流
	defer jsonFile.Close()
	if err != nil {
		return nil
	}
	//把json文件转化为map
	messageMap := make(map[string]string)
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&messageMap)
	if err != nil {
		return nil
	}
	return messageMap
}

//获取app目录
func GetAppPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	filepath := dir + "\\app"
	return filepath
}
