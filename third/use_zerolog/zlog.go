package zlog

import (
	"os"

	"github.com/rs/zerolog"
)

func ZTest() {
	file, err := os.OpenFile("test.log",
		os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log := zerolog.New(file).With().Timestamp().Logger()

	log.Print("hello world")
}
