package pkgfoo

import (
	"fmt"

	zl "github.com/krls08/go-with-tests/600_logger/zlog/logger2"
)

type IFoo interface {
	Hello()
}

type defFoo struct {
	name string
}

func NewDefFoo(n string) *defFoo {
	zl.Zlog.Info().Msg("hello")
	return &defFoo{
		name: n,
	}
}

func (f defFoo) Hello() {
	fmt.Println("hello from fmt", f.name)
	zl.Zlog.Info().Msg("test from zero log")
	//zlog.Print("hello")
	//l := logger.Get()
	//l.Print("hello")
}
