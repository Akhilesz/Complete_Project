package multiplication

import (
	"context"
	"main/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func MultiplyViaGRPC(val int64, times int64) (res int64, err error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	client := proto.NewAdditionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var currResult int64 = 0
	for i := 0; int64(i) < times; i++ {
		resp, err := client.Add(ctx, &proto.AddRequest{
			A: currResult,
			B: val,
		})
		if err != nil {
			return 0, err
		}
		res += resp.Res
	}
	return res, nil
}
