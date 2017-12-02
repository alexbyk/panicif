package panicif_test

import (
	"fmt"
	"testing"

	"github.com/alexbyk/ftest"
	"github.com/alexbyk/panicif"
)

func Test_Err(t *testing.T) {
	ft := ftest.New(t)
	ft.PanicsSubstr(func() { panicif.Err(fmt.Errorf("Foo")) }, "Foo")
}
