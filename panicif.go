/*Package panicif provides shorthands for error handling. Instead of:

  if err != nil {
    panic(err)
  }

you can write:
  panicif.Err(err)

Be aware that an invocation to Err may not be inlined (last time I checked it's not), so
don't use it for a critical piece of code
*/
package panicif

import (
	"fmt"
	"reflect"
)

// Err panics with a given argument if the argument isn't nil
func Err(err error) {
	if err != nil {
		panic(err)
	}
}

// True panics if the first argument is true
func True(cond bool, format string, a ...interface{}) {
	if cond {
		panic(fmt.Errorf(format, a...))
	}
}

// False panics if the first argument is true
func False(cond bool, format string, a ...interface{}) {
	if !cond {
		panic(fmt.Errorf(format, a...))
	}
}

// Nil panics if the first argument is Nil
func Nil(val interface{}, format string, a ...interface{}) {
	if isNil(val) {
		panic(fmt.Errorf(format, a...))
	}
}

// NotNil panics if the first argument is Nil
func NotNil(val interface{}, format string, a ...interface{}) {
	if !isNil(val) {
		panic(fmt.Errorf(format, a...))
	}
}

func isNil(v interface{}) bool {
	return v == nil || isNilable(v) && reflect.ValueOf(v).IsNil()
}

func isNilable(v interface{}) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return true
	}
	return false
}
