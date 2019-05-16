package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type creative struct {
	ig *Instagram
}

func newCreative(i *Instagram) *creative {
	return &creative{ig: i}
}

func (c *creative) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = c.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
