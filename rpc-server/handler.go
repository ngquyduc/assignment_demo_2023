package main

import (
	"context"
	"github.com/ngquyduc/assignment_demo_2023/rpc-server/db"
	"github.com/ngquyduc/assignment_demo_2023/rpc-server/kitex_gen/rpc"
	"log"
	"time"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	req.Message.SendTime = time.Now().Unix()
	resp := rpc.NewSendResponse()
	err := db.LoadMsg(req.Message)
	if err != nil {
		resp.Code = 500
		resp.Msg = "Failed"
		log.Fatalln(err)
		return resp, err
	}
	resp.Code = 0
	resp.Msg = "Success"
	return resp, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	resp := rpc.NewPullResponse()
	if req.Reverse == nil {
		req.Reverse = new(bool)
		*req.Reverse = false
	}
	messages, hasMore, nextCursor, err := db.GetMsg(req.Chat, req.Cursor, int64(req.Limit), *req.Reverse)
	if err != nil {
		resp.Code = 500
		resp.Msg = "Failed"
		log.Fatalln(err)
		return resp, err
	}
	resp.Code = 0
	resp.Msg = "Success"
	resp.Messages = messages
	resp.HasMore = &hasMore
	resp.NextCursor = &nextCursor

	return resp, nil
}
