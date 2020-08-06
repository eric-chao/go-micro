package handler

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/store"

	log "github.com/micro/go-micro/v2/logger"

	post "post/proto/post"
)

type PostHandler struct{
	Store  store.Store
	Client client.Client
}

func (h *PostHandler) Save(ctx context.Context, req *post.SaveRequest, rsp *post.SaveResponse) error {
	log.Info("Received Posts.Save request")
	return nil
}
