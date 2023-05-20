package main

import (
	zl "github.com/krls08/go-with-tests/600_logger/zlog/logger2"
	"github.com/krls08/go-with-tests/600_logger/zlog/pkgfoo"
)

func main() {
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//zlog.Print("hello world")
	//zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	//zlog.Info().Msg("Info message")
	//zlog.Error().Msg("Error message")
	//file, err := os.OpenFile(
	//	"myapp.log",
	//	os.O_APPEND|os.O_CREATE|os.O_WRONLY,
	//	0664,
	//)
	//if err != nil {
	//	panic(err)
	//}

	//defer file.Close()

	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//logger := zerolog.New(file).With().Timestamp().Logger()

	zl.Init()

	//l := logger.Get()

	//l.Info().Msg("Info message")

	foo := pkgfoo.NewDefFoo("krls")

	foo.Hello()
}
