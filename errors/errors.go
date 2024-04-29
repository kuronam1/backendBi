package errors

import (
	"log"
	"log/slog"
)

func HandleSimpleErr(e error) {
	if e != nil {
		panic(e)
	}
}

func HandleInitErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Error(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
