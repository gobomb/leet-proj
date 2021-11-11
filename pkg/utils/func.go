package utils

import (
	"reflect"
	"runtime"
	"strings"
)

func FuncName(f interface{}) string {
	sp := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	return sp[len(sp)-1]
}
