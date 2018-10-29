package handler

import (
	"context"
	"fmt"
	"os"

	pb "gitlab.com/gitlab-org/gitaly-proto/go"
	"gitlab.com/gitlab-org/gitaly/client"
)

func ReceivePack(gitalyAddress string, request *pb.SSHReceivePackRequest) (int32, error) {
	if gitalyAddress == "" {
		return 0, fmt.Errorf("no gitaly_address given")
	}

	gitalyAddress, isTls := transFormTls(gitalyAddress)
	conn, err := client.Dial(gitalyAddress, dialOpts(isTls))

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return client.ReceivePack(ctx, conn, os.Stdin, os.Stdout, os.Stderr, request)
}
