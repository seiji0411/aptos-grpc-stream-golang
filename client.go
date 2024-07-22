package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"aptos-grpc-stream-golang/grpcurl"

	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

var maxMsgSz = 30000000
var target = "grpc.mainnet.aptoslabs.com:443"
var data = "{ \"starting_version\": 1044414248 }"
var symbol = "aptos.indexer.v1.RawData/GetTransactions"
var headers = []string{"authorization:Bearer aptoslabs_JAcBRSFCvNL_9PPj4k9jYWRSLDZyzDHqD9bKv9qWoMSmm"}

func fail(err error, msg string, args ...interface{}) {
	if err != nil {
		msg += ": %v"
		args = append(args, err)
	}
	fmt.Fprintf(os.Stderr, msg, args...)
	fmt.Fprintln(os.Stderr)
	if err != nil {
		os.Exit(1)
	} else {
		// nil error means it was CLI usage issue
		fmt.Fprintf(os.Stderr, "Try '%s -help' for more details.\n", os.Args[0])
		os.Exit(2)
	}
}

type CustomHandler struct {
}

func (_h *CustomHandler) Write(p []byte) (n int, err error) {
	fmt.Printf("Transaction length: %s\n", p)
	return len(p), nil
}

func NewCustomHandler() io.Writer {
	return &CustomHandler{}
}

func main() {
	ctx := context.Background()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSz)))

	var creds credentials.TransportCredentials
	tlsConf, err := grpcurl.ClientTLSConfig(false, "", "", "")
	if err != nil {
		fail(err, "Failed to create TLS config")
	}
	creds = credentials.NewTLS(tlsConf)

	cc, err := grpcurl.BlockingDial(ctx, "tcp", target, creds, opts...)
	if err != nil {
		fail(err, "Failed to dial target host %q", target)
	}

	in := strings.NewReader(data)

	var descSource grpcurl.DescriptorSource
	md := grpcurl.MetadataFromHeaders(headers)
	refCtx := metadata.NewOutgoingContext(ctx, md)

	refClient := grpcreflect.NewClientAuto(refCtx, cc)
	refClient.AllowMissingFileDescriptors()
	reflSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)
	descSource = reflSource

	// if not verbose output, then also include record delimiters
	// between each message, so output could potentially be piped
	// to another grpcurl process
	options := grpcurl.FormatOptions{
		EmitJSONDefaultFields: false,
		IncludeTextSeparator:  true,
		AllowUnknownFields:    false,
	}
	rf, formatter, err := grpcurl.RequestParserAndFormatter(grpcurl.Format("json"), descSource, in, options)
	if err != nil {
		fail(err, "Failed to construct request parser and formatter for %q", "json")
	}

	h := &grpcurl.DefaultEventHandler{
		Out:            NewCustomHandler(),
		Formatter:      formatter,
		VerbosityLevel: 0,
	}

	err = grpcurl.InvokeRPC(ctx, descSource, cc, symbol, headers, h, rf.Next)
	if err != nil {
		fail(err, "Error invoking method %q", symbol)
	}
}
