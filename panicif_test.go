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

func Test_True(t *testing.T) {
	ft := ftest.New(t)
	ft.PanicsSubstr(func() { panicif.True(true, "Foo") }, "Foo")
	ft.PanicsSubstr(func() { panicif.True(true, "Foo %s", "bar") }, "Foo bar")
}

func Test_False(t *testing.T) {
	ft := ftest.New(t)
	ft.PanicsSubstr(func() { panicif.False(false, "Foo") }, "Foo")
	ft.PanicsSubstr(func() { panicif.False(false, "Foo %s", "bar") }, "Foo bar")
}
