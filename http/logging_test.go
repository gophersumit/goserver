package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	. "github.com/contiamo/goserver/http"
)

var _ = Describe("Logging", func() {
	It("should be possible to configure logging", func() {
		buf := &bytes.Buffer{}
		logrus.SetOutput(buf)
		srv, err := createServer([]Option{WithLogging("test")})
		Expect(err).NotTo(HaveOccurred())
		ts := httptest.NewServer(srv.Handler)
		_, err = http.Get(ts.URL + "/logging")
		Expect(err).NotTo(HaveOccurred())
		Expect(strings.Contains(buf.String(), "completed handling request")).To(BeTrue())
	})
})
