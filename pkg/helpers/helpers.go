package helpers

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// Empty 类似于 PHP 的 empty() 函数
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

func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// FirstElement 安全地获取 args[0]，避免 panic: runtime error: index out of range
func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b{
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RemoveEmptyToString(a []string, seps... string) string {
	var sep string
	if len(sep) > 0 {
		sep = seps[0]
	} else {
		sep = "/"
	}
	var ret []string
	for _, v := range a{
		if v == "" {
			continue
		}
		ret = append(ret, v)
	}
	return strings.Join(ret, "/")
}

func RemoveEmptyToArray(a string, seps... string) []string {
	var sep string
	if len(sep) > 0 {
		sep = seps[0]
	} else {
		sep = "/"
	}
	var ret []string
	array := strings.Split(a, sep)
	for _, v := range array{
		if v == "" {
			continue
		}
		ret = append(ret, v)
	}
	return ret
}
