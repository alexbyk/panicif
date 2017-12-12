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
