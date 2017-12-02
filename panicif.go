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

// Err panics with a given argument if the argument isn't nil
func Err(err error) {
	if err != nil {
		panic(err)
	}
}
