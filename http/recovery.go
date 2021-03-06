package http

import (
	"io"
	"net/http"

	recovery "github.com/bakins/net-http-recover"
)

// WithRecovery configures panic recovery for that server
func WithRecovery(writer io.Writer, printStack bool) Option {
	return &recoveryOption{writer, printStack}
}

type recoveryOption struct {
	writer     io.Writer
	printStack bool
}

func (opt *recoveryOption) WrapHandler(handler http.Handler) http.Handler {
	return recovery.Handler(opt.writer, handler, opt.printStack)
}
