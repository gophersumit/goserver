package server_test

import (
	"bytes"
	"context"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/contiamo/goserver/grpc"
	"github.com/contiamo/goserver/grpc/test"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

var _ = Describe("Logging", func() {
	It("should be possible to setup logging option", func() {
		buf := &bytes.Buffer{}
		logrus.SetOutput(buf)
		logrus.SetLevel(logrus.DebugLevel)
		srv, err := createServerWithOptions([]Option{WithMDLogging()})
		Expect(err).NotTo(HaveOccurred())
		Expect(srv).NotTo(BeNil())
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go ListenAndServe(ctx, ":3003", srv)
		cli, err := createPlaintextTestClient(ctx, "localhost:3003")
		Expect(err).NotTo(HaveOccurred())

		md := metadata.New(map[string]string{"test": "value"})
		ctx = metadata.NewOutgoingContext(ctx, md)

		resp, err := cli.Ping(ctx, &test.PingReq{Msg: "test"})
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Msg).To(Equal("test"))
		Expect(strings.Contains(buf.String(), "finished unary call with code OK")).To(BeTrue())
		Expect(strings.Contains(buf.String(), "test:\"value\"")).To(BeTrue())
	})
})